<template>
  <div>chat room

    <div>
      <button @click="send">发消息</button>
    </div>
  </div>
</template>

<script>
export default { name: 'ChatRoom',
  data () {
    return {
      websock: null
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
      const wsuri = 'ws://127.0.0.1:9999/api/ws?username=test'
      this.websock = new WebSocket(wsuri)
      this.websock.onopen = this.open
      this.websock.onclose = this.close
      this.websock.onmessage = this.receive
        .console.log(this.websock)
      // this.websock.send('hhh')
    },
    send: function () {
      let data = '123'
      console.log('send : ', data)
      this.websock.send(data)
    },
    open: function () {
      console.log('socket连接成功')
    },
    close: function () {
      console.log('连接中断')
    },
    receive: function (e) {
      console.log('receive : ', e)
    }
  }
}
</script>
