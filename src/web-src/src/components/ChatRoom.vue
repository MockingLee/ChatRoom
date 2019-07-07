<template>
  <div>chat room

    <div class="container">

      <!-- <ul>
        <li v-for="item in message_array"
            :key='item'>{{item}}</li>
      </ul>
      <input type="text"
             v-model="message" />
      <button @click="send">add</button>
    </div>
    <div>
      <button>发消息</button>
    </div> -->
    <kendo-chat ref="chat"
                @post="post"
                v-on:post="onPost"
                v-on:sendmessage="onSendMessage"
                v-on:actionclick="onActionClick"
                v-on:typingstart="onTypingStart"
                v-on:typingend="onTypingEnd"></kendo-chat>

  </div>
  </div>
</template>

<script>
import '../global'

export default { name: 'ChatRoom',
  data () {
    return {
      websock: null
      // message_array: [],
      // message: ''
    }
  },
  created () {
    this.username = this.$route.params.username
    this.initWebSocket()
  },
  // destroyed () {
  //   this.websock.close()
  // },
  methods: {
    post (args) {
      console.log(args)
    },
    initWebSocket () {
      var username = this.username
      const wsuri = 'ws://' + global.socketHost + '/api/ws?username=' + username
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
      this.handleMsg(e)
    },
    onPost: function (ev) {
      console.log('A message has been posted to the Chat widget!')
    },
    onSendMessage: function (ev) {
      // console.log(ev.text)
      console.log('A message has been posted to the Chat widget using the message box!')
      this.websock.send(ev.text)
    },
    onActionClick: function (ev) {
      console.log('The user clicked an action button in attachment template, or selected a suggestedAction!')
    },
    onTypingStart: function (ev) {
      console.log('The user started typing in the Chat message box!')
    },
    onTypingEnd: function (ev) {
      console.log('The user cleared the Chat message box!')
    },
    handleMsg (msg) {
      // msg : json
      console.log(msg)
      console.log(msg.Type)
      if (msg.Type === global.Event.Message && msg.User !== this.username) {
        this.$refs.chat.kendoWidget().renderMessage({
          type: 'text',
          text: msg.Content
        }, {
          id: msg.User,
          name: msg.User
        })
      }
      if (msg.Type === global.Event.Join) {
        this.$refs.chat.kendoWidget().renderMessage({
          type: 'text',
          text: msg.Content
        }, {
          id: msg.User,
          name: msg.User
        })
      }
      // this.$refs.chat.kendoWidget().renderMessage({
      //   type: msg.,
      //   text: 'Hello Kendo Chat'
      // }, {
      //   id: '123',
      //   name: 'Sample User',
      //   iconUrl: 'https://demos.telerik.com/kendo-ui/content/web/chat/avatar.png'
      // })
    }
  }
}
</script>
