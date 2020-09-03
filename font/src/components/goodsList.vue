<template>
  <div class="clearfix">
    <section class="hami-price-plan-area mt-50">
      <div class="container">
        <div class="row">
          <div class="col-12">
            <div class="section-heading text-center">
              <h2>商品列表</h2>
              <p>
                您还可获取专有分享链接，享受60%交易提成
                <Poptip title="专有分享链接 - 通过sealyun赚钱" placement="right-end">
                  <div slot="content" class="poptip">
                    <p>{{ shareLink }}</p>
                    <p>任何用户通过上面链接访问网站并成功交易您将获得交易的60%提成</p>
                    <p>如嵌入自己的markdown文档中, 发到群里，或者把链接直接发给有需要的朋友</p>
                  </div>
                  <span id="sharelink">推广链接</span>
                </Poptip>
              </p>
            </div>
          </div>
        </div>
        <div class="clearfix" style="margin-bottom:20px">
          <div class="tabs">
            <div v-for="(item,index ) in tabs" :key="index" @click="tabF(item.name)" :class="item.name === tab ? 'tab-active' : ''">{{item.txt}}</div>
          </div>
        </div>
        <div v-if="tab === 'allGoods'">
          <div class="Box">
            <div class="Box-header">
              <h3 class="Box-title">
                <svg class="octicon octicon-file-zip" viewBox="0 0 16 16" version="1.1" width="16" height="16" aria-hidden="true">
                  <path fill-rule="evenodd" d="M3.5 1.75a.25.25 0 01.25-.25h3a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h2.086a.25.25 0 01.177.073l2.914 2.914a.25.25 0 01.073.177v8.586a.25.25 0 01-.25.25h-.5a.75.75 0 000 1.5h.5A1.75 1.75 0 0014 13.25V4.664c0-.464-.184-.909-.513-1.237L10.573.513A1.75 1.75 0 009.336 0H3.75A1.75 1.75 0 002 1.75v11.5c0 .649.353 1.214.874 1.515a.75.75 0 10.752-1.298.25.25 0 01-.126-.217V1.75zM8.75 3a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM6 5.25a.75.75 0 01.75-.75h.5a.75.75 0 010 1.5h-.5A.75.75 0 016 5.25zm2 1.5A.75.75 0 018.75 6h.5a.75.75 0 010 1.5h-.5A.75.75 0 018 6.75zm-1.25.75a.75.75 0 000 1.5h.5a.75.75 0 000-1.5h-.5zM8 9.75A.75.75 0 018.75 9h.5a.75.75 0 010 1.5h-.5A.75.75 0 018 9.75zm-.75.75a1.75 1.75 0 00-1.75 1.75v3c0 .414.336.75.75.75h2.5a.75.75 0 00.75-.75v-3a1.75 1.75 0 00-1.75-1.75h-.5zM7 12.25a.25.25 0 01.25-.25h.5a.25.25 0 01.25.25v2.25H7v-2.25z"></path>
                </svg>
                离线安装包
              </h3>
            </div>
            <div class="Box-row position-relative" v-for="(p,index) in products" :key="index">
              <div class="flex-auto min-width-0">
                <div class="commit js-details-container Details">
                  <div>
                    <h4 class="flex-auto min-width-0 pr-2 pb-1 commit-title">
                      {{ p.name }}
                      <span style="color: #586069!important;font-size:12px">￥{{ p.price }}</span>
                    </h4>

                    <Tooltip content="购买完自动跳转下载，已付款点击自动下载">
                      <a class="buy-btn" :href="p.url" target="_blank" rel="noopener">
                        购买
                      </a>
                    </Tooltip>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div v-if="tab === 'selfGoods'">

        </div>

      </div>
    </section>
    <Back-top></Back-top>
  </div>
</template>

<script  type="text/javascript">
import VueCookies from 'vue-cookies'
export default {
  props: {
    msg: String,
  },
  methods: {
    ok() {
      // this.$Message.info('Clicked ok')
    },
  },
  data: function () {
    var d = {
      time: 'http://store.lameleg.com:8080/pro/kubernetes1.13.1?time=' + new Date().getTime(),
      shareLink: '',
      share: false,
      products: [],
      tab: 'allGoods',
      tabs: [
        {
          name: 'allGoods',
          txt: '所有商品',
        },
        {
          name: 'selfGoods',
          txt: '已购商品',
        },
      ],
    }
    if (typeof this.$route.query.referrer != 'undefined') {
      VueCookies.set('referrer', this.$route.query.referrer)
    }
    if (VueCookies.get('referrer') != null) {
      d.time += '&referrer=' + VueCookies.get('referrer')
    }

    this.$http
      .get('http://store.lameleg.com:8080/loginless/info/user', {
        credentials: true,
      })
      .then(function (res) {
        if (typeof res.data.login != 'undefined') {
          d.shareLink = 'http://store.lameleg.com?referrer=' + res.data.login
        } else {
          d.shareLink = '登录后才能看到您的推广连接'
        }
      })

    this.$http
      .get('http://store.lameleg.com:8080/loginless/pro', { credentials: true })
      .then(function (res) {
        for (var i = 0; i < res.data.length; i++) {
          var p = {}
          p.name = res.data[i].ProductName
          p.url =
            'http://store.lameleg.com:8080/pro/' +
            p.name +
            '?time=' +
            new Date().getTime() +
            '&referrer=' +
            VueCookies.get('referrer')
          p.price = res.data[i].ProductPrice
          d.products.push(p)
        }
        d.products.reverse()
      })
    return d
  },
  methods: {
    tabF(tab) {
      this.tab = tab
    },
  },
}
</script>
 <style scoped>
.poptip p {
  font-size: 13px;
  text-align: left;
}
#sharelink {
  cursor: pointer;
  color: #2c66de;
  font-size: 13px;
}
.tabs {
  font-size: 14px;
  float: left;
  margin: 0px auto 0;
  box-shadow: 0 0 20p 1px rgba(0, 0, 0, 0.1);
  border-radius: 5px;
  overflow: hidden;
}
.tabs div {
  width: 100px;
  float: left;
  text-align: center;
  padding: 8px 16px;
  font-weight: 500;
  line-height: 24px;
  cursor: pointer;
}
.tab-active {
  color: #fff;
  background-color: #0366d6;
  border-color: #005cc5;
}

.Box {
  background-color: #fff;
  border: 1px solid #e1e4e8;
  border-radius: 6px;
  max-height: 899px;
  overflow: scroll;
  margin-bottom: 50px;
}

.Box-header {
  padding: 16px;
  margin: -1px -1px 0;
  background-color: #f6f8fa;
  border: 1px solid #e1e4e8;
  border-top-left-radius: 6px;
  border-top-right-radius: 6px;
}
.Box-title {
  font-size: 14px;
  font-weight: 600;
  margin-bottom: 0;
}
.Box-row {
  padding: 16px 16px 10px;
  margin-top: -1px;
  list-style-type: none;
  border-top: 1px solid #e1e4e8;
}
.buy-btn {
  font-size: 14px;
  border-radius: 5px;
  background-color: #0366d6;
  padding: 3px 10px;
  color: #fff;
}
.ivu-tooltip {
  position: absolute;
  right: 10px;
  top: 17px;
}
.commit-title {
  color: #0366d6;
  font-weight: 600px;
}
.pb-1,
.py-1 {
  padding-bottom: 0;
}
</style>