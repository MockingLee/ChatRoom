<template>
  <div>chat room

    <div class="container">

      <ul>
        <li v-for="item in message_array"
            :key='item'>{{item}}</li>
      </ul>
      <input type="text"
             v-model="message" />
      <button @click="send">add</button>
    </div>
    <div>
      <button>发消息</button>
    </div>
  </div>
</template>

<script>
export default { name: 'ChatRoom',
  data () {
    return {
      websock: null,
      message_array: [],
      message: ''
    }
  },
  created () {
    this.initWebSocket()
  },
  destroyed () {
    this.websock.close()
  },
  methods: {

    initWebSocket () {
      var username = this.$route.params.username
      const wsuri = 'ws://127.0.0.1:9999/api/ws?username=' + username
      this.websock = new WebSocket(wsuri)
      this.websock.onopen = this.open
      this.websock.onclose = this.close
      this.websock.onmessage = this.receive
      console.log(this.websock)
      // this.websock.send('hhh')
    },
    send: function () {
      console.log('send : ', this.message)
      this.websock.send(this.message)
    },
    open: function () {
      console.log('socket连接成功')
    },
    close: function () {
      console.log('连接中断')
    },
    receive: function (e) {
      console.log('receive : ', e)
      e = JSON.parse(e.data)
      console.log(e.Type)
      if (e.Type === 2) { this.message_array.push(e.User + ' says ' + e.Content) }
      if (e.Type === 1) { this.message_array.push(e.User + ' quit ') }
      if (e.Type === 0) { this.message_array.push(e.User + ' in ') }
      console.log(this.message_array)
      // this.message_array.push('123')
    }
  }
}
</script>
