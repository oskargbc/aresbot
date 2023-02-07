package aio

import (
	"aresbot/internal/aio/cli"
	"aresbot/internal/aio/cli/utils"
	"aresbot/internal/aio/constants"
	"aresbot/internal/aio/discord"
	"aresbot/internal/aio/global"
	models2 "aresbot/internal/aio/models"
	shared2 "aresbot/internal/aio/shared"
	"aresbot/internal/aio/shops"
	"aresbot/internal/aio/tools"
	"aresbot/pkg/clear"
	l "aresbot/pkg/logger"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/hugolgst/rich-go/client"
	"github.com/kardianos/osext"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func NewBot() MainBot {
	return MainBot{}
}

type MainBot struct {
}

func (b *MainBot) Run() {
	clear.CallClear()
	groups, settings, proxys := ReadInfo()
	version := "1.0.0"

	go b.DiscordPrecense(version)

	data := checkLicense(settings.License)

	w := shared2.WebhookHandler{
		HyperData: data,
		Version:   version,
		PrivatUrl: settings.WebhookURL,
		PublicUrl: "https://discord.com/api/webhooks/922073808913977404/4oLhWAvHVo8I1BbRVK5SK8UcVFaWXQ2ML1veq9FxksrDfdN70nWI7Z05Dyb4DN9jIwyr",
	}

	shop := []string{}
	for s, _ := range shops.ShopBots {
		shop = append(shop, s)
	}
	tool := []string{}
	for t, _ := range tools.AllTools {
		tool = append(tool, t)
	}

	cliData := shared2.CliData{
		Hyper:          data,
		Version:        version,
		Groups:         groups,
		Settings:       settings,
		Proxys:         proxys,
		WebhookHandler: w,
		Shops:          shop,
		Tools:          tool,
	}

	c := cli.NewCli(cliData)
	start := "Menu"

	// main loop
	for {
		resp := c.Handler(&start)

		switch *resp.Cmd {
		case constants.CmdReload:
			groups, settings, proxys = ReadInfo()
			c.Data.Groups = groups
			c.Data.Proxys = proxys
			c.Data.Settings = settings
			time.Sleep(time.Second * 1)
			utils.PrintWithTag("MENU", "Reloaded data")
		case constants.CmdToggleQuicktasks:
			if constants.WithQuickTask {
				constants.WithQuickTask = false
			} else {
				constants.WithQuickTask = true
			}
		case constants.CmdRunTasks:
			r := shops.BotRunner{
				Tasks:          resp.Tasks,
				Proxys:         proxys,
				Settings:       settings,
				WebhookHandler: w,
			}
			r.Run()

			utils.PrintWithTag("MENU", "[B] Go back to menu")

			again := true
			for again {
				input := utils.AwaitInput("MENU", "Choose action: ", []string{})
				if input == nil {
					continue
				}
				again = false
				if *input == "B" {
					start = "Menu"
				}
			}
		case constants.CmdStartTool:
			t := tools.ToolRunner{
				Tasks:          resp.Tasks,
				Proxys:         proxys,
				Settings:       settings,
				WebhookHandler: w,
			}
			t.Run(resp.ToolName, resp.Extras)

			utils.PrintWithTag("MENU", "[B] Go back to menu")

			again := true
			for again {
				input := utils.AwaitInput("MENU", "Choose action: ", []string{})
				if input == nil {
					continue
				}
				again = false
				if *input == "B" {
					start = "Menu"
				}
			}
		case constants.CmdStartDiscordMonitor:
			d := discord.NewDiscordMonitor(settings.DiscordChannelIds, settings.DiscordServerIds, settings.DiscordKeywords, settings.DiscordToken)

			d.Run()
		}
	}
}

func (b *MainBot) DiscordPrecense(v string) {

	err := client.Login("936376271385006101")
	if err != nil {
		l.ErrorLogger.Println(err, "Couldn't connect CLient Precense")
	}

	t := time.Now()

	err = client.SetActivity(client.Activity{
		State:      "Ares " + v,
		Details:    "godly cooking",
		LargeImage: "",
		LargeText:  "Ares",
		SmallImage: "",
		SmallText:  "destroying everthing",
		Timestamps: &client.Timestamps{
			Start: &t,
		},
		Buttons: []*client.Button{
			&client.Button{
				Label: "Dashboard",
				Url:   "https://aresbot.hyper.co/dashboard",
			},
		},
	})
	if err != nil {
		l.ErrorLogger.Println(err, "Couldn't connect Precense")
	}
}

func formatProxy(elem string, settings models2.Settings) url.URL {
	protocol := "http"

	if settings.ProxyHttps {
		protocol = "https"
	}

	stringParts := strings.Split(elem, ":")

	if len(stringParts) == 2 {
		elem = fmt.Sprintf("%s://%s:%s", protocol, stringParts[0], stringParts[1])
	}

	if len(stringParts) == 4 {
		elem = fmt.Sprintf("%s://%s:%s@%s:%s", protocol, stringParts[2], stringParts[3], stringParts[0], stringParts[1])
	}

	// windows bug case replace
	if strings.Contains(elem, "\r") {
		elem = strings.ReplaceAll(elem, "\r", "")
	}

	url, err := url.Parse(elem)

	if err != nil {
		l.ErrorLogger.Println(err, "Failed Parsing Proxy")
	}

	return *url
}

func fetchVersion() string {
	shared2.OInfoF("Checking version")

	version := constants.CurrentVersion

	resp, err := http.Get("https://firestormbot.com/api/version/check")
	if err != nil {
		l.ErrorLogger.Println(err, " checking version")
		os.Exit(-1)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var Response shared2.ApiVersionResponse
	err = json.Unmarshal(body, &Response)
	if err != nil {
		l.ErrorLogger.Println(err, " checking version")
		os.Exit(-1)
	}

	if version == Response.Version {
		shared2.OInfoF("Version: " + version + " everything up to date")
		time.Sleep(time.Millisecond * 500)
		return version
	} else {
		shared2.OErrorF("Current: " + version + ", New: " + Response.Version + ", please update soon")
		time.Sleep(time.Millisecond * 500)
		return version
	}

	return version
}

func checkLicense(license string) shared2.HyperResponse {
	c := http.Client{}
	var response shared2.HyperResponse
	var resp *http.Response

	shared2.OInfoF("Checking License")
	req, err := http.NewRequest("GET", "https://api.hyper.co/v4/licenses/"+license, nil)
	if err != nil {
		l.ErrorLogger.Println(err)
		shared2.OErrorF("Failed checking License")
		time.Sleep(3 * time.Second)
		os.Exit(-1)
	}

	req.Header.Add("Authorization", "Bearer sk_klU24iBE6qsRqluIqdjyqkzyl9rmgeW0YsC5QtyXBXnfnehUAG1mipPJvzzr5Zit")

	resp, err = c.Do(req)
	if err != nil {
		l.ErrorLogger.Println(err)
		shared2.OErrorF("Check Failed")
		time.Sleep(3 * time.Second)
		os.Exit(-1)
	}

	time.Sleep(time.Second * 1)

	if resp.StatusCode == 200 {
		shared2.OInfoF("Logged in")

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			l.ErrorLogger.Println(err)
			shared2.OErrorF("Data reading failed")
			time.Sleep(3 * time.Second)
			os.Exit(-1)
		}

		err = json.Unmarshal(body, &response)
		if err != nil {
			l.ErrorLogger.Println(err)
			shared2.OErrorF("Data reading failed")
			time.Sleep(3 * time.Second)
			os.Exit(-1)
		}

		return response
	}

	shared2.OErrorF("Invalid License")
	time.Sleep(3 * time.Second)
	os.Exit(-1)
	return response
}

func ReadInfo() ([]models2.TaskGroup, models2.Settings, []models2.Proxy) {
	var profiles []models2.Profile
	tasksGroups := []models2.TaskGroup{}
	var settings models2.Settings
	var proxys []models2.Proxy

	// Read all info
	executablePath, _ := osext.ExecutableFolder()

	os.Chdir(executablePath)

	file, err := ioutil.ReadFile("profile.csv")
	if err != nil {
		l.ErrorLogger.Println(err)

		fmt.Println("Couldn't find profile.csv, closing..")

		time.Sleep(5 * time.Second)
		os.Exit(-1)
	}
	err = gocsv.Unmarshal(bytes.NewReader(file), &profiles)
	if err != nil {
		l.ErrorLogger.Println(err)

		fmt.Println("Error in file profile.csv, closing.. ")

		time.Sleep(5 * time.Second)
		fmt.Println(err)
	}

	/*
		file, err = ioutil.ReadFile("task.json")
		if err != nil {
			l.ErrorLogger.Println(err)

			fmt.Println("Couldn't find task.json, closing..")

			time.Sleep(5 * time.Second)
			os.Exit(-1)
		}
		err = json.Unmarshal(file, &tasks)
		if err != nil {
			l.ErrorLogger.Println(err)

			fmt.Println("Error in file task.json, closing.. ")

			time.Sleep(5 * time.Second)
			fmt.Println(err)
		}
	*/
	file, err = ioutil.ReadFile("settings.json")
	if err != nil {
		l.ErrorLogger.Println(err)

		fmt.Println("Couldn't find settings.json, closing..")

		time.Sleep(5 * time.Second)
		os.Exit(-1)
	}
	err = json.Unmarshal(file, &settings)
	if err != nil {
		l.ErrorLogger.Println(err)

		fmt.Println("Error in file settings.json, closing.. ")

		time.Sleep(5 * time.Second)
		fmt.Println(err)
	}

	f, _ := ioutil.ReadDir("./tasks/")
	for _, f := range f {
		if strings.Contains(f.Name(), "tasks_") {
			tasks := []models2.Task{}
			file, err = ioutil.ReadFile("./tasks/" + f.Name())

			if err != nil {
				l.ErrorLogger.Println(err)

				fmt.Println("Couldn't find " + f.Name() + ", closing..")

				time.Sleep(5 * time.Second)
				os.Exit(-1)
			}
			err = gocsv.Unmarshal(bytes.NewReader(file), &tasks)
			if err != nil {
				l.ErrorLogger.Println(err)

				fmt.Println("Error in file " + f.Name() + ", closing.. ")

				time.Sleep(5 * time.Second)
				fmt.Println(err)
			}

			newtasks := tasks
			for k, task := range tasks {
				task.Mode = strings.ToUpper(task.Mode)
				for _, profil := range profiles {
					if profil.ProfileName == task.ProfileName {
						newtasks[k].Profile = profil
					}
					if profil.ProfileName == settings.QuicktaskProfile {
						global.SettingQuicktaskProfile = profil
					}
				}
			}

			group := models2.TaskGroup{
				Name:  f.Name(),
				Tasks: tasks,
			}

			tasksGroups = append(tasksGroups, group)
		}
	}

	file, err = ioutil.ReadFile("proxys.txt")
	if err != nil {
		l.ErrorLogger.Println(err)

		fmt.Println("Couldn't find proxys.txt, closing..")

		time.Sleep(5 * time.Second)
		os.Exit(-1)
	}
	strFile := string(file)
	txtProxyList := strings.Split(strFile, "\n")
	if len(txtProxyList) > 0 {
		for _, txtproxy := range txtProxyList {
			p := formatProxy(txtproxy, settings)

			proxys = append(proxys, models2.Proxy{Url: p})
		}
	}

	//Set up tasks
	global.SettingQuicktaskAmount = settings.QuicktaskAmount
	global.SettingQuicktaskProfileName = settings.QuicktaskProfile
	global.SettingQuicktaskWithProxy = settings.QuicktaskWithProxy

	shared2.OInfoF("Loaded Files")

	return tasksGroups, settings, proxys
}
