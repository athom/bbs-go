<template>
  <div>
    <section class="main">
      <div class="container">
        <div class="left-main-container">
          <div class="m-left">
            <div class="topic-detail topic-wrap">
              <div class="header">
                <div class="left">
                  <div
                    class="avatar avatar-size-45 is-rounded"
                    :style="{backgroundImage:'url('+ topic.user.avatar +')'}"
                  />
                </div>
                <div class="center">
                  <a :href="'/topic/' + topic.topicId" :title="topic.title">
                    <div class="topic-title">{{ topic.title }}</div>
                  </a>

                  <div class="topic-meta">
                    <span class="meta-item">
                      <a :href="'/user/' + topic.user.id">{{ topic.user.nickname }}</a>
                    </span>
                    <span class="meta-item">
                      {{ topic.lastCommentTime | prettyDate }}
                    </span>
                    <span class="meta-item">
                      <span v-for="tag in topic.tags" :key="tag.tagId" class="tag">
                        <a :href="'/topics/tag/' + tag.tagId + '/1'">{{ tag.tagName }}</a>
                      </span>
                    </span>
                    <span class="meta-item act">
                      <a @click="addFavorite(topic.topicId)">
                        <i class="iconfont icon-favorite" />&nbsp;{{ favorited ? '已收藏' : '收藏' }}
                      </a>
                    </span>
                    <span v-if="isOwner" class="meta-item act">
                      <a @click="deleteTopic(topic.topicId)">
                        <i class="iconfont icon-delete" />&nbsp;删除
                      </a>
                    </span>
                    <span v-if="isOwner" class="meta-item act">
                      <a :href="'/topic/edit/' + topic.topicId">
                        <i class="iconfont icon-edit" />&nbsp;修改
                      </a>
                    </span>
                  </div>
                </div>
                <div class="right">
                  <div class="like">
                    <span class="like-btn" :class="{'liked': topic.liked}" @click="like(topic)">
                      <i class="iconfont icon-like" />
                    </span>
                    <span v-if="topic.likeCount" class="like-count">{{ topic.likeCount }}</span>
                  </div>
                  <span class="count">{{ topic.commentCount }}&nbsp;/&nbsp;{{ topic.viewCount }}</span>
                </div>
              </div>

              <div class="content" v-html="topic.content" />

              <ins
                class="adsbygoogle"
                style="display:block"
                data-ad-format="fluid"
                data-ad-layout-key="-ig-s+1x-t-q"
                data-ad-client="ca-pub-5683711753850351"
                data-ad-slot="4728140043"
              />
              <script>
                (adsbygoogle = window.adsbygoogle || []).push({});
              </script>
            </div>

            <!-- 评论 -->
            <comment entity-type="topic" :entity-id="topic.topicId" :show-ad="true" />
          </div>
          <div class="m-right">
            <!-- 展示广告190x90 -->
            <ins
              class="adsbygoogle"
              style="display:inline-block;width:190px;height:90px"
              data-ad-client="ca-pub-5683711753850351"
              data-ad-slot="9345305153"
            />
            <script>
              (adsbygoogle = window.adsbygoogle || []).push({});
            </script>

            <!-- 展示广告190x190 -->
            <ins
              class="adsbygoogle"
              style="display:inline-block;width:190px;height:190px"
              data-ad-client="ca-pub-5683711753850351"
              data-ad-slot="5685455263"
            />
            <script>
              (adsbygoogle = window.adsbygoogle || []).push({});
            </script>

            <div v-if="topic.toc" ref="toc" class="toc widget">
              <div class="header">
                目录
              </div>
              <div class="content" v-html="topic.toc" />
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import utils from '~/common/utils'
import Comment from '~/components/Comment'
export default {
  components: {
    Comment
  },
  computed: {
    isOwner: function () {
      return (
        this.currentUser &&
        this.topic &&
        this.currentUser.id === this.topic.user.id
      )
    }
  },
  async asyncData({ $axios, params, error }) {
    let topic
    try {
      topic = await $axios.get('/api/topic/' + params.id)
    } catch (e) {
      error({
        statusCode: 404,
        message: '话题不存在'
      })
      return
    }

    const currentUser = await $axios.get('/api/user/current')
    const favorited = await $axios.get('/api/favorite/favorited', {
      params: {
        entityType: 'topic',
        entityId: params.id
      }
    })

    return {
      currentUser: currentUser,
      topic: topic,
      favorited: favorited.favorited
    }
  },
  mounted() {
    utils.handleToc()
  },
  head() {
    return {
      title: this.$siteTitle(this.topic.title)
    }
  },
  methods: {
    async addFavorite(topicId) {
      try {
        if (this.favorited) {
          await this.$axios.get('/api/favorite/delete', {
            params: {
              entityType: 'topic',
              entityId: topicId
            }
          })
          this.favorited = false
          this.$toast.success('已取消收藏！')
        } else {
          await this.$axios.get('/api/topic/favorite/' + topicId)
          this.favorited = true
          this.$toast.success('收藏成功')
        }
      } catch (e) {
        console.error(e)
        this.$toast.error('收藏失败：' + (e.message || e))
      }
    },
    async deleteTopic(topicId) {
      try {
        await this.$axios.post('/api/topic/delete/' + topicId)
        this.$toast.success('删除成功', {
          duration: 2000,
          onComplete: function () {
            utils.linkTo('/topics')
          }
        })
      } catch (e) {
        console.error(e)
        this.$toast.error('删除失败：' + (e.message || e))
      }
    },
    async like(topic) {
      try {
        await this.$axios.get('/api/topic/like/' + topic.topicId)
        topic.liked = true
        topic.likeCount++
      } catch (e) {
        if (e.errorCode === 1) {
          this.$toast.info('请登录后点赞！！！', {
            action: {
              text: '去登录',
              onClick: (e, toastObject) => {
                utils.toSignin()
              }
            }
          })
        } else {
          this.$toast.error(e.message || e)
        }
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.aside-avatar {
  width: 150px;
  height: 150px;
}
.topic-detail {
  margin-bottom: 20px;

  .content {
    padding-top: 10px;
    font-size: 15px;
    color: #000;
  }
}
</style>
