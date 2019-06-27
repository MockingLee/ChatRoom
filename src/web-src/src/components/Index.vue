<template>
  <div>
    <el-row :gutter="20">
      <el-col :span="6">
        <div class="grid-content bg-purple"></div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple">
          <el-form :model="formInline"
                   class="demo-form-inline">
            <el-form-item label="用户名">
              <el-input v-model="formInline.user"
                        placeholder="用户名"></el-input>
            </el-form-item>
            <el-form-item>
              <el-button type="primary"
                         @click="onSubmit('formInline')">登录</el-button>
            </el-form-item>
          </el-form>

        </div>
      </el-col>
      <el-col :span="6">
        <div class="grid-content bg-purple"></div>
      </el-col>
    </el-row>

  </div>

</template>

<script>
import api from '../axios'

export default {
  name: 'Index',

  data () {
    return {
      formInline: {
        user: ''
      }
    }
  },
  methods: {
    onSubmit () {
      console.log(this.formInline.user)

      api.userIn(this.formInline.user).then(({
        data
      }) => {
        console.log(data)
        data = JSON.parse(data)
        console.log(data.success)
        if (data.success === 'true') {
          this.$router.push({
            name: 'ChatRoom',
            params: {
              'username': this.formInline.user
            }
          })
        } else if (data.success === 'false') alert(data.msg)
      })

      // this.$router.push({
      //   name: 'ChatRoom',
      //   params: {
      //     'username': this.formInline.user
      //   }
      // })

      // let opt = JSON.stringify({
      //   'username': this.formInline.user
      // })
      // console.log(opt)
      // api.getTest(opt).then(({
      //   data
      // }) => {
      //   data = JSON.parse(data)
      //   if (data.success) {
      //     console.log('ok')
      //     this.$router.push({
      //       name: 'ChatRoom',
      //       params: {
      //         'username': this.formInline.user
      //       }
      //     })
      //   }
      // })
    }
  }

}
</script>
