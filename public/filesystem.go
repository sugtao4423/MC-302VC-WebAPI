// Code generated by shogo82148/assets-life v1.0.1. DO NOT EDIT.

//go:generate go run assets-life.go "../assets" . public

package public

import (
	"io"
	"net/http"
	"os"
	"path"
	"sort"
	"strings"
	"time"
)

// Root is the root of the file system.
var Root http.FileSystem = fileSystem{
	file{
		name:    "/",
		content: "",
		mode:    0755 | os.ModeDir,
		next:    0,
		child:   1,
	},
	file{
		name:    "/index.html",
		content: "<!DOCTYPE html>\n<html>\n  <head>\n    <title>MC-302VC-WebAPI</title>\n    <meta charset=\"utf-8\" />\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\" />\n\n    <link\n      href=\"https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css\"\n      rel=\"stylesheet\"\n      integrity=\"sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC\"\n      crossorigin=\"anonymous\"\n    />\n    <style>\n      [v-cloak] {\n        display: none;\n      }\n      .buttons {\n        display: flex;\n        align-items: end;\n        margin-left: -0.8rem;\n      }\n      .buttons > * {\n        margin-top: 0.8rem;\n        margin-left: 0.8rem;\n      }\n      .bath-auto-timer-setting {\n        width: 20rem;\n      }\n    </style>\n  </head>\n\n  <body>\n    <div id=\"app\" class=\"container py-4\" v-cloak>\n      <div class=\"position-fixed top-0 end-0 p-3\">\n        <div\n          id=\"liveToast\"\n          class=\"toast text-white bg-danger\"\n          :class=\"{ show: error.show, hide: !error.show }\"\n          role=\"alert\"\n          aria-live=\"assertive\"\n          aria-atomic=\"true\"\n        >\n          <div class=\"toast-body\">{{error.message}}</div>\n        </div>\n      </div>\n\n      <div v-if=\"!isLoggedIn\">\n        <form>\n          <div class=\"row justify-content-center mb-3\">\n            <div class=\"col-md-6\">\n              <label for=\"usernameForm\" class=\"form-label\">Username</label>\n              <input\n                v-model=\"username\"\n                type=\"text\"\n                class=\"form-control\"\n                id=\"usernameForm\"\n              />\n            </div>\n          </div>\n\n          <div class=\"row justify-content-center mb-3\">\n            <div class=\"col-md-6\">\n              <label for=\"passwordForm\" class=\"form-label\">Password</label>\n              <input\n                v-model=\"password\"\n                type=\"password\"\n                class=\"form-control\"\n                id=\"passwordForm\"\n              />\n            </div>\n          </div>\n\n          <div class=\"row justify-content-center text-end\">\n            <div class=\"col-md-6\">\n              <button type=\"button\" class=\"btn btn-primary\" @click=\"login\">\n                Login\n              </button>\n            </div>\n          </div>\n        </form>\n      </div>\n\n      <div v-if=\"data !== null\">\n        <div>給湯機動作状態: {{readableStatus.operationStatus}}</div>\n        <div>給湯温度: {{readableStatus.waterTemp}}</div>\n        <div>風呂温度: {{readableStatus.bathTemp}}</div>\n        <div>風呂自動タイマー: {{readableStatus.bathAutoTimerStatus}}</div>\n        <div>風呂自動タイマー時間: {{readableStatus.bathAutoTimerTime}}</div>\n        <div>風呂動作状態: {{readableStatus.bathOperationStatus}}</div>\n        <div>風呂自動: {{readableStatus.bathAutoModeStatus}}</div>\n        <div>風呂追い焚き: {{readableStatus.bathAdditionalHeatingStatus}}</div>\n\n        <div class=\"mt-4 buttons\">\n          <button\n            type=\"button\"\n            :class=\"toggleButtonClass.bathAutoTimerStatus\"\n            @click=\"toggleBathAutoTimer\"\n          >\n            風呂自動タイマー{{toggleString.bathAutoTimerStatus}}\n          </button>\n          <div class=\"input-group bath-auto-timer-setting\">\n            <input\n              v-model=\"bathAutoTime\"\n              type=\"time\"\n              class=\"form-control\"\n              aria-describedby=\"bath-timer\"\n            />\n            <button\n              class=\"btn btn-success\"\n              type=\"button\"\n              id=\"bath-timer\"\n              @click=\"setBathAutoTimerTime\"\n            >\n              設定\n            </button>\n          </div>\n        </div>\n        <div class=\"buttons\">\n          <button\n            type=\"button\"\n            :class=\"toggleButtonClass.bathAutoModeStatus\"\n            @click=\"toggleBathAutoMode\"\n          >\n            風呂自動{{toggleString.bathAutoModeStatus}}\n          </button>\n          <button\n            type=\"button\"\n            :class=\"toggleButtonClass.bathAdditionalHeatingStatus\"\n            @click=\"toggleBathAdditionalHeating\"\n          >\n            風呂追い焚き{{toggleString.bathAdditionalHeatingStatus}}\n          </button>\n        </div>\n      </div>\n    </div>\n\n    <script src=\"https://unpkg.com/vue@3/dist/vue.global.prod.js\"></script>\n    <script src=\"./main.js\"></script>\n  </body>\n</html>\n",
		mode:    0644,
		next:    2,
		child:   -1,
	},
	file{
		name:    "/main.js",
		content: "const KeyBasic = 'basic'\n\nconst initLoggedIn = () => {\n  const basic = sessionStorage.getItem(KeyBasic)\n  return basic !== null && basic !== ''\n}\n\nVue.createApp({\n  data() {\n    return {\n      apiBaseUrl: './api',\n      username: '',\n      password: '',\n      basic: sessionStorage.getItem(KeyBasic) ?? '',\n      data: null,\n      bathAutoTime: '',\n      isLoggedIn: initLoggedIn(),\n      error: {\n        show: false,\n        message: '',\n      },\n    }\n  },\n  mounted() {\n    if (this.isLoggedIn) {\n      this.getData()\n    }\n  },\n  methods: {\n    async login() {\n      this.basic = btoa(`${this.username}:${this.password}`)\n      await this.getData()\n      sessionStorage.setItem(KeyBasic, this.basic)\n    },\n    async getData() {\n      const res = await fetch(`${this.apiBaseUrl}/status`, {\n        headers: { Authorization: `Basic ${this.basic}` },\n      })\n      if (!res.ok) {\n        this.showError('Server Error')\n        this.isLoggedIn = false\n        return\n      }\n      const data = await res.json()\n      if (data.error) {\n        this.showError('Data Error')\n        this.isLoggedIn = false\n        return\n      }\n      this.data = data\n      this.isLoggedIn = true\n    },\n    async postApi(path, json) {\n      const res = await fetch(this.apiBaseUrl + path, {\n        method: 'POST',\n        headers: {\n          Authorization: `Basic ${this.basic}`,\n          'Content-Type': 'application/json',\n        },\n        body: JSON.stringify(json),\n      })\n      if (!res.ok) {\n        this.showError('Server Error')\n        return\n      }\n      const data = await res.json()\n      if (data.error) {\n        this.showError('Data Error')\n        return\n      }\n\n      setTimeout(() => this.getData(), 1000)\n    },\n    showError(msg) {\n      this.error = {\n        show: true,\n        message: msg,\n      }\n      setTimeout(() => {\n        this.error.show = false\n      }, 3000)\n    },\n    async toggleBathAutoTimer() {\n      await this.postApi('/bathAutoTimer', {\n        status: !this.data.bath_auto_timer_status,\n      })\n    },\n    async setBathAutoTimerTime() {\n      const [hour, minute] = this.bathAutoTime.split(':')\n      if (Number.isNaN(hour) || Number.isNaN(minute)) {\n        this.showError('Invalid Time')\n        return\n      }\n\n      await this.postApi('/bathAutoTimer/time', {\n        hour: Number(hour),\n        minute: Number(minute),\n      })\n    },\n    async toggleBathAutoMode() {\n      await this.postApi('/bath/auto', {\n        status: !this.data.bath_auto_mode_status,\n      })\n    },\n    async toggleBathAdditionalHeating() {\n      await this.postApi('/bath/additionalHeating', {\n        status: !this.data.bath_additional_heating_status,\n      })\n    },\n  },\n  computed: {\n    readableStatus() {\n      const onoff = (bool) => (bool ? 'ON' : 'OFF')\n      return {\n        operationStatus: onoff(this.data.operation_status),\n        waterTemp: this.data.water_temperature + '℃',\n        bathTemp: this.data.bath_temperature + '℃',\n        bathAutoTimerStatus: onoff(this.data.bath_auto_timer_status),\n        bathAutoTimerTime: this.data.bath_auto_timer_time\n          .map((n) => (n < 10 ? `0${n}` : n))\n          .join(':'),\n        bathOperationStatus: onoff(this.data.bath_operation_status),\n        bathAutoModeStatus: onoff(this.data.bath_auto_mode_status),\n        bathAdditionalHeatingStatus: onoff(\n          this.data.bath_additional_heating_status\n        ),\n      }\n    },\n    toggleButtonClass() {\n      const classes = (bool) => 'btn ' + (bool ? 'btn-warning' : 'btn-info')\n      return {\n        bathAutoTimerStatus: classes(this.data.bath_auto_timer_status),\n        bathAutoModeStatus: classes(this.data.bath_auto_mode_status),\n        bathAdditionalHeatingStatus: classes(\n          this.data.bath_additional_heating_status\n        ),\n      }\n    },\n    toggleString() {\n      const str = (bool) => (bool ? 'OFF' : 'ON')\n      return {\n        bathAutoTimerStatus: str(this.data.bath_auto_timer_status),\n        bathAutoModeStatus: str(this.data.bath_auto_mode_status),\n        bathAdditionalHeatingStatus: str(\n          this.data.bath_additional_heating_status\n        ),\n      }\n    },\n  },\n}).mount('#app')\n",
		mode:    0644,
		next:    -1,
		child:   -1,
	},
}

type fileSystem []file

func (fs fileSystem) Open(name string) (http.File, error) {
	name = path.Clean("/" + name)
	i := sort.Search(len(fs), func(i int) bool { return fs[i].name >= name })
	if i >= len(fs) || fs[i].name != name {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
	}
	f := &fs[i]
	return &httpFile{
		Reader: strings.NewReader(f.content),
		file:   f,
		fs:     fs,
		idx:    i,
		dirIdx: f.child,
	}, nil
}

type file struct {
	name    string
	content string
	mode    os.FileMode
	child   int
	next    int
}

var _ os.FileInfo = (*file)(nil)

func (f *file) Name() string {
	return path.Base(f.name)
}

func (f *file) Size() int64 {
	return int64(len(f.content))
}

func (f *file) Mode() os.FileMode {
	return f.mode
}

var zeroTime time.Time

func (f *file) ModTime() time.Time {
	return zeroTime
}

func (f *file) IsDir() bool {
	return f.Mode().IsDir()
}

func (f *file) Sys() interface{} {
	return nil
}

type httpFile struct {
	*strings.Reader
	file   *file
	fs     fileSystem
	idx    int
	dirIdx int
}

var _ http.File = (*httpFile)(nil)

func (f *httpFile) Stat() (os.FileInfo, error) {
	return f.file, nil
}

func (f *httpFile) Readdir(count int) ([]os.FileInfo, error) {
	ret := []os.FileInfo{}
	if !f.file.IsDir() {
		return ret, nil
	}

	if count <= 0 {
		for f.dirIdx >= 0 {
			entry := &f.fs[f.dirIdx]
			ret = append(ret, entry)
			f.dirIdx = entry.next
		}
		return ret, nil
	}

	ret = make([]os.FileInfo, 0, count)
	for f.dirIdx >= 0 {
		entry := &f.fs[f.dirIdx]
		ret = append(ret, entry)
		f.dirIdx = entry.next
		if len(ret) == count {
			return ret, nil
		}
	}
	return ret, io.EOF
}

func (f *httpFile) Close() error {
	return nil
}