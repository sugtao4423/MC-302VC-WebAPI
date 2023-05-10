const KeyBasic = 'basic'

const initLoggedIn = () => {
  const basic = sessionStorage.getItem(KeyBasic)
  return basic !== null && basic !== ''
}

Vue.createApp({
  data() {
    return {
      apiBaseUrl: './api',
      username: '',
      password: '',
      basic: sessionStorage.getItem(KeyBasic) ?? '',
      data: null,
      bathAutoTime: '',
      isLoggedIn: initLoggedIn(),
      error: {
        show: false,
        message: '',
      },
    }
  },
  mounted() {
    if (this.isLoggedIn) {
      this.getData()
    }
  },
  methods: {
    async login() {
      this.basic = btoa(`${this.username}:${this.password}`)
      await this.getData()
      sessionStorage.setItem(KeyBasic, this.basic)
    },
    async getData() {
      const res = await fetch(`${this.apiBaseUrl}/status`, {
        headers: { Authorization: `Basic ${this.basic}` },
      })
      if (!res.ok) {
        this.showError('Server Error')
        this.isLoggedIn = false
        return
      }
      const data = await res.json()
      if (data.error) {
        this.showError('Data Error')
        this.isLoggedIn = false
        return
      }
      this.data = data
      this.isLoggedIn = true
    },
    async postApi(path, json) {
      const res = await fetch(this.apiBaseUrl + path, {
        method: 'POST',
        headers: {
          Authorization: `Basic ${this.basic}`,
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(json),
      })
      if (!res.ok) {
        this.showError('Server Error')
        return
      }
      const data = await res.json()
      if (data.error) {
        this.showError('Data Error')
        return
      }

      setTimeout(() => this.getData(), 1000)
    },
    showError(msg) {
      this.error = {
        show: true,
        message: msg,
      }
      setTimeout(() => {
        this.error.show = false
      }, 3000)
    },
    async toggleBathAutoTimer() {
      await this.postApi('/bathAutoTimer', {
        status: !this.data.bath_auto_timer_status,
      })
    },
    async setBathAutoTimerTime() {
      const [hour, minute] = this.bathAutoTime.split(':')
      if (Number.isNaN(hour) || Number.isNaN(minute)) {
        this.showError('Invalid Time')
        return
      }

      await this.postApi('/bathAutoTimer/time', {
        hour: Number(hour),
        minute: Number(minute),
      })
    },
    async toggleBathAutoMode() {
      await this.postApi('/bath/auto', {
        status: !this.data.bath_auto_mode_status,
      })
    },
    async toggleBathAdditionalHeating() {
      await this.postApi('/bath/additionalHeating', {
        status: !this.data.bath_additional_heating_status,
      })
    },
  },
  computed: {
    readableStatus() {
      const onoff = (bool) => (bool ? 'ON' : 'OFF')
      return {
        operationStatus: onoff(this.data.operation_status),
        waterTemp: this.data.water_temperature + '℃',
        bathTemp: this.data.bath_temperature + '℃',
        bathAutoTimerStatus: onoff(this.data.bath_auto_timer_status),
        bathAutoTimerTime: this.data.bath_auto_timer_time
          .map((n) => (n < 10 ? `0${n}` : n))
          .join(':'),
        bathOperationStatus: onoff(this.data.bath_operation_status),
        bathAutoModeStatus: onoff(this.data.bath_auto_mode_status),
        bathAdditionalHeatingStatus: onoff(
          this.data.bath_additional_heating_status
        ),
      }
    },
    toggleButtonClass() {
      const classes = (bool) => 'btn ' + (bool ? 'btn-warning' : 'btn-info')
      return {
        bathAutoTimerStatus: classes(this.data.bath_auto_timer_status),
        bathAutoModeStatus: classes(this.data.bath_auto_mode_status),
        bathAdditionalHeatingStatus: classes(
          this.data.bath_additional_heating_status
        ),
      }
    },
    toggleString() {
      const str = (bool) => (bool ? 'OFF' : 'ON')
      return {
        bathAutoTimerStatus: str(this.data.bath_auto_timer_status),
        bathAutoModeStatus: str(this.data.bath_auto_mode_status),
        bathAdditionalHeatingStatus: str(
          this.data.bath_additional_heating_status
        ),
      }
    },
  },
}).mount('#app')
