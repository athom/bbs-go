<template>
  <section class="main">
    <div class="container">
      <div class="columns">
        <div class="column is-9">
          <div class="main-body">
            <div class="widget">
              <div class="header">
                <nav class="breadcrumb" aria-label="breadcrumbs" style="margin-bottom: 0px;">
                  <ul>
                    <li>
                      <a href="/">首页</a>
                    </li>
                    <li>
                      <a :href="'/user/' + user.id">{{ user.nickname }}</a>
                    </li>
                    <li class="is-active">
                      <a href="#" aria-current="page">编辑资料</a>
                    </li>
                  </ul>
                </nav>
              </div>
              <div class="content">
                <div class="field is-horizontal">
                  <div class="field-label is-normal">
                    <label class="label">用户名：</label>
                  </div>
                  <div class="field-body">
                    <div class="field">
                      <div class="control has-icons-left">
                        <label v-if="user.username">{{ user.username }}</label>
                        <a v-else @click="showSetUsername = true">点击设置</a>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="field is-horizontal">
                  <div class="field-label is-normal">
                    <label class="label">邮箱：</label>
                  </div>
                  <div class="field-body">
                    <div class="field">
                      <div class="control has-icons-left">
                        <label v-if="user.email">{{ user.email }}</label>
                        <a v-else @click="showSetEmail = true">点击设置</a>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="field is-horizontal">
                  <div class="field-label is-normal">
                    <label class="label">密码：</label>
                  </div>
                  <div class="field-body">
                    <div class="field">
                      <div class="control has-icons-left">
                        <template v-if="user.passwordSet">
                          <label>密码已设置&nbsp;</label>
                          <a @click="showUpdatePassword = true">点击修改</a>
                        </template>
                        <a v-else @click="showSetPassword = true">点击设置</a>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="field is-horizontal">
                  <div class="field-label is-normal">
                    <label class="label">
                      <span style="color:red;">*&nbsp;</span>昵称：
                    </label>
                  </div>
                  <div class="field-body">
                    <div class="field">
                      <div class="control has-icons-left">
                        <input
                          v-model="user.nickname"
                          name="nickname"
                          class="input is-success"
                          type="text"
                          placeholder="请输入昵称"
                        >
                        <span class="icon is-small is-left">
                          <i class="iconfont icon-username" />
                        </span>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="field is-horizontal">
                  <div class="field-label is-normal">
                    <label class="label">
                      <span style="color:red;">*&nbsp;</span>头像：
                    </label>
                  </div>
                  <div class="field-body">
                    <div class="field">
                      <div class="control">
                        <img :src="user.avatar" style="width: 150px;height:150px;">
                        <div class="file">
                          <label class="file-label">
                            <input class="file-input" type="file" @change="uploadAvatar">
                            <span class="file-cta">
                              <span class="file-icon">
                                <i class="iconfont icon-upload" />
                              </span>
                              <span class="file-label">选择头像</span>
                            </span>
                          </label>
                        </div>
                        <span style="font-weight: bold; color:red;">*图像必须为正方形，大小不要超过1M。</span>
                      </div>
                    </div>
                  </div>
                </div>

                <div class="field is-horizontal">
                  <div class="field-label is-normal">
                    <label class="label">简介：</label>
                  </div>
                  <div class="field-body">
                    <div class="field">
                      <div class="control">
                        <textarea
                          v-model="user.description"
                          name="description"
                          class="textarea"
                          rows="2"
                          placeholder="一句话介绍你自己"
                        />
                      </div>
                    </div>
                  </div>
                </div>

                <div class="field is-horizontal">
                  <div class="field-label is-normal" />
                  <div class="field-body">
                    <div class="field">
                      <div class="control">
                        <a class="button is-success" @click="submitForm">提交修改</a>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div class="column is-3">
          <div class="main-aside">
            <user-center-sidebar :user="user" :current-user="user" />
          </div>
        </div>
      </div>
    </div>

    <!-- 设置用户名 -->
    <div class="modal" :class="{'is-active': showSetUsername}">
      <div class="modal-background" />
      <div class="modal-card">
        <div class="widget">
          <div class="header">
            设置用户名
            <button class="delete" aria-label="close" @click="showSetUsername = false" />
          </div>
          <div class="content">
            <div class="field">
              <div class="control has-icons-left">
                <input
                  v-model="username"
                  class="input is-success"
                  type="text"
                  placeholder="请输入用户名"
                  @keydown.enter="setUsername"
                >
                <span class="icon is-small is-left">
                  <i class="iconfont icon-username" />
                </span>
              </div>
            </div>
          </div>
          <div class="footer is-right">
            <a class="button is-success" @click="setUsername">确定</a>
            <a class="button" @click="showSetUsername = false">取消</a>
          </div>
        </div>
      </div>
    </div>

    <!-- 设置邮箱 -->
    <div class="modal" :class="{'is-active': showSetEmail}">
      <div class="modal-background" />
      <div class="modal-card">
        <div class="widget">
          <div class="header">
            设置邮箱
            <button class="delete" aria-label="close" @click="showSetEmail = false" />
          </div>
          <div class="content">
            <div class="field">
              <div class="control has-icons-left">
                <input
                  v-model="email"
                  class="input is-success"
                  type="text"
                  placeholder="请输入邮箱"
                  @keydown.enter="setEmail"
                >
                <span class="icon is-small is-left">
                  <i class="iconfont icon-email" />
                </span>
              </div>
            </div>
          </div>
          <div class="footer is-right">
            <a class="button is-success" @click="setEmail">确定</a>
            <a class="button" @click="showSetEmail = false">取消</a>
          </div>
        </div>
      </div>
    </div>

    <!-- 设置密码 -->
    <div class="modal" :class="{'is-active': showSetPassword}">
      <div class="modal-background" />
      <div class="modal-card">
        <div class="widget">
          <div class="header">
            设置密码
            <button class="delete" aria-label="close" @click="showSetPassword = false" />
          </div>
          <div class="content">
            <div class="field">
              <div class="control has-icons-left">
                <input
                  v-model="password"
                  class="input is-success"
                  type="password"
                  placeholder="请输入密码"
                  @keydown.enter="setPassword"
                >
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>
            <div class="field">
              <div class="control has-icons-left">
                <input
                  v-model="rePassword"
                  class="input is-success"
                  type="password"
                  placeholder="请再次确认密码"
                  @keydown.enter="setPassword"
                >
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>
          </div>
          <div class="footer is-right">
            <a class="button is-success" @click="setPassword">确定</a>
            <a class="button" @click="showSetPassword = false">取消</a>
          </div>
        </div>
      </div>
    </div>

    <!-- 修改密码 -->
    <div class="modal" :class="{'is-active': showUpdatePassword}">
      <div class="modal-background" />
      <div class="modal-card">
        <div class="widget">
          <div class="header">
            修改密码
            <button class="delete" aria-label="close" @click="showUpdatePassword = false" />
          </div>
          <div class="content">
            <div class="field">
              <div class="control has-icons-left">
                <input
                  v-model="oldPassword"
                  class="input is-success"
                  type="password"
                  placeholder="请输入当前密码"
                  @keydown.enter="updatePassword"
                >
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>
            <div class="field">
              <div class="control has-icons-left">
                <input
                  v-model="password"
                  class="input is-success"
                  type="password"
                  placeholder="请输入密码"
                  @keydown.enter="updatePassword"
                >
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>
            <div class="field">
              <div class="control has-icons-left">
                <input
                  v-model="rePassword"
                  class="input is-success"
                  type="password"
                  placeholder="请再次确认密码"
                  @keydown.enter="updatePassword"
                >
                <span class="icon is-small is-left">
                  <i class="iconfont icon-password" />
                </span>
              </div>
            </div>
          </div>
          <div class="footer is-right">
            <a class="button is-success" @click="updatePassword">确定</a>
            <a class="button" @click="showUpdatePassword = false">取消</a>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script>
import UserCenterSidebar from '~/components/UserCenterSidebar'
export default {
  middleware: 'authenticated',
  components: {
    UserCenterSidebar
  },
  data() {
    return {
      showSetUsername: false,
      username: '',
      showSetEmail: false,
      email: '',
      showSetPassword: false, // 显示设置密码
      showUpdatePassword: false, // 显示修改密码
      password: '',
      rePassword: '',
      oldPassword: ''
    }
  },
  head() {
    return {
      title: this.$siteTitle(this.user.nickname + ' - 编辑资料')
    }
  },
  async asyncData({ $axios, params }) {
    const [user] = await Promise.all([$axios.get('/api/user/current')])
    return {
      user: user
    }
  },
  methods: {
    async submitForm() {
      try {
        await this.$axios.post('/api/user/edit/' + this.user.id, {
          nickname: this.user.nickname,
          avatar: this.user.avatar,
          description: this.user.description
        })
        this.$toast.success('修改成功')
      } catch (e) {
        console.error(e)
        this.$toast.error('修改失败：' + (e.message || e))
      }
    },
    async uploadAvatar(e) {
      const files = e.target.files
      if (files.length <= 0) {
        return
      }
      try {
        const file = files[0]
        const formData = new FormData()
        formData.append('image', file, file.name)
        const ret = await this.$axios.post('/api/upload', formData, {
          headers: { 'Content-Type': 'multipart/form-data' }
        })
        this.user.avatar = ret.url
      } catch (e) {
        console.error(e)
      }
    },
    async setUsername() {
      try {
        const me = this
        await this.$axios.post('/api/user/set/username', {
          username: me.username
        })
        this.user = await this.$axios.get('/api/user/current')
        this.$toast.success('设置成功')
        this.username = ''
        this.showSetUsername = false
      } catch (err) {
        this.$toast.error('设置失败：' + (err.message || err))
      }
    },
    async setEmail() {
      try {
        const me = this
        await this.$axios.post('/api/user/set/email', {
          email: me.email
        })
        this.user = await this.$axios.get('/api/user/current')
        this.$toast.success('设置成功')
        this.email = ''
        this.showSetEmail = false
      } catch (err) {
        this.$toast.error('设置失败：' + (err.message || err))
      }
    },
    async setPassword() {
      try {
        const me = this
        await this.$axios.post('/api/user/set/password', {
          password: me.password,
          rePassword: me.rePassword
        })
        this.user = await this.$axios.get('/api/user/current')
        this.$toast.success('设置成功')
        this.password = ''
        this.rePassword = ''
        this.showSetPassword = false
      } catch (err) {
        this.$toast.error('设置失败：' + (err.message || err))
      }
    },
    async updatePassword() {
      try {
        const me = this
        await this.$axios.post('/api/user/update/password', {
          oldPassword: me.oldPassword,
          password: me.password,
          rePassword: me.rePassword
        })
        this.user = await this.$axios.get('/api/user/current')
        this.$toast.success('修改成功')
        this.password = ''
        this.rePassword = ''
        this.oldPassword = ''
        this.showUpdatePassword = false
      } catch (err) {
        this.$toast.error('修改失败：' + (err.message || err))
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.control {
  a,
  label {
    // padding-top: .375em;
    // font-size: 14px;
    line-height: 32px;
  }
}

.modal {
  .widget {
    background: #ffffff;
    margin: 0px;
    padding: 10px;
  }
}
</style>
