<template>
  <footer class="footer-area section-padding-80-0">
    <div class="main-footer-area">
      <div class="container">
        <div class="row justify-content-between">
          <div class="col-8 col-sm-5 col-lg-2">
            <div class="single-footer-widget mb-80">
              <h5 class="widget-title">联系方式</h5>
              <ul class="footer-nav">
                <li><a href="https://sealyun.com/" target="_blank">官网</a></li>
                <li><a href="http://store.lameleg.com" target="_blank">官方商城</a></li>
                <li><a href="https://t.me/gsealyun" target="_blank">Telegram</a> </li>
                <li>QQ群：98488045</li>
                <li>钉钉群:35371178 </li>
                <li>邮箱：fhtjob@hotmail.com</li>
              </ul>
            </div>
          </div>

          <div class="col-5 col-sm-4 col-lg-2">
            <div class="single-footer-widget mb-80">
              <h5 class="widget-title">开源项目</h5>
              <ul class="footer-nav">
                <li><a href="https://github.com/fanux/sealos" target="_blank">kubernetes一键HA</a></li>
                <li><a href="https://github.com/fanux/fist" target="_blank">轻量级kubernetes管理工具</a></li>
                <li><a href="https://github.com/fanux/lhttp" target="_blank">好用的websocket框架</a></li>
              </ul>
            </div>
          </div>

          <div class="col-5 col-sm-3 col-lg-2">
            <div class="single-footer-widget mb-80">
              <h5 class="widget-title">友情链接</h5>

              <ul class="footer-nav">
                <li><a href="https://kuboard.cn/#sealyun" target="_blank">kuboard k8s dashboard</a></li>
                <li><a href="https://www.yangcs.net#sealyun" target="_blank">骚客米开朗琪杨</a></li>
                <li><a href="https://www.qikqiak.com/?utm_source=sealyun.com" target="_blank">阳明的博客</a></li>
                <li><a href="https://zhangguanzhang.github.io/#sealyun" target="_blank">张馆长</a></li>
              </ul>
            </div>
          </div>

          <div class="col-12 col-sm-8 col-lg-4">
            <div class="single-footer-widget mb-80">
              <h5 class="widget-title">公众号</h5>
              <img src="https://sealyun.com/kubernetes-qrcode.jpg" width="135">
            </div>
          </div>

        </div>
      </div>
    </div>

    <div class="bottom-footer-area bg-gray">
      <div class="container">
        <div class="row align-items-center">
          <div>
            <div class="copywrite-text">
              <p>
                备案信息。。。。。 by <a href="#" target="_blank">fanux</a>
              </p>
            </div>
          </div>
          <!-- <div style="width:320px">
            <div class="copywrite-text">
              <p>
                <a href="https://github.com/fanux/sealos" target="_blank">sealyun</a> 专注于kubernetes安装 <i class="fa fa-heart-o" aria-hidden="true"></i> by <a href="#" target="_blank">fanux</a>
              </p>
            </div>
          </div>
        <div style="flex:1">
            <div class="payments-methods d-flex align-items-center">
              <p>核心贡献者</p>
              <a v-for="(item,index) in developer" :key="index" :title="item.author.path.slice(1)" :href="'https://github.com'+item.author.path"><img :src="item.author.avatar" alt=""></a>
            </div>
          </div> -->

        </div>
      </div>
    </div>
  </footer>
</template>

<script type="text/javascript">
export default {
  name: 'footCommon',
  data() {
    return {
      developer: [],
    }
  },
  created() {
    this.getDeveloper()
  },
  methods: {
    getDeveloper() {
      let developer = [
        {
          total: 1,
          author: {
            id: 7947984,
            login: 'wuxingzhong',
            avatar: 'https://avatars0.githubusercontent.com/u/7947984?s=60&v=4',
            path: '/wuxingzhong',
            hovercard_url: '/users/wuxingzhong/hovercard',
          },
        },
        {
          total: 1,
          author: {
            id: 16590391,
            login: 'vimerr',
            avatar: 'https://avatars1.githubusercontent.com/u/16590391?s=60&v=4',
            path: '/vimerr',
            hovercard_url: '/users/vimerr/hovercard',
          },
        },
        {
          total: 1,
          author: {
            id: 261698,
            login: 'ideal',
            avatar: 'https://avatars3.githubusercontent.com/u/261698?s=60&v=4',
            path: '/ideal',
            hovercard_url: '/users/ideal/hovercard',
          },
        },
        {
          total: 1,
          author: {
            id: 12635393,
            login: 'fzkun',
            avatar: 'https://avatars2.githubusercontent.com/u/12635393?s=60&v=4',
            path: '/fzkun',
            hovercard_url: '/users/fzkun/hovercard',
          },
        },
        {
          total: 1,
          author: {
            id: 12584881,
            login: 'dujiacheng',
            avatar: 'https://avatars0.githubusercontent.com/u/12584881?s=60&v=4',
            path: '/dujiacheng',
            hovercard_url: '/users/dujiacheng/hovercard',
          },
        },
        {
          total: 1,
          author: {
            id: 31763211,
            login: 'aaronlyc',
            avatar: 'https://avatars0.githubusercontent.com/u/31763211?s=60&v=4',
            path: '/aaronlyc',
            hovercard_url: '/users/aaronlyc/hovercard',
          },
        },
        {
          total: 2,
          author: {
            id: 18624174,
            login: 'mxfj',
            avatar: 'https://avatars1.githubusercontent.com/u/18624174?s=60&v=4',
            path: '/mxfj',
            hovercard_url: '/users/mxfj/hovercard',
          },
        },
        {
          total: 2,
          author: {
            id: 39077997,
            login: 'fengkuangdestone',
            avatar: 'https://avatars1.githubusercontent.com/u/39077997?s=60&v=4',
            path: '/fengkuangdestone',
            hovercard_url: '/users/fengkuangdestone/hovercard',
          },
        },
        {
          total: 2,
          author: {
            id: 31427850,
            login: 'ImgBotApp',
            avatar: 'https://avatars1.githubusercontent.com/u/31427850?s=60&v=4',
            path: '/ImgBotApp',
            hovercard_url: '/users/ImgBotApp/hovercard',
          },
        },
        {
          total: 3,
          author: {
            id: 18641678,
            login: 'zhangguanzhang',
            avatar: 'https://avatars1.githubusercontent.com/u/18641678?s=60&v=4',
            path: '/zhangguanzhang',
            hovercard_url: '/users/zhangguanzhang/hovercard',
          },
        },
        {
          total: 3,
          author: {
            id: 4159479,
            login: 'svolence',
            avatar: 'https://avatars3.githubusercontent.com/u/4159479?s=60&v=4',
            path: '/svolence',
            hovercard_url: '/users/svolence/hovercard',
          },
        },
        {
          total: 3,
          author: {
            id: 30151093,
            login: 'LucienShui',
            avatar: 'https://avatars0.githubusercontent.com/u/30151093?s=60&v=4',
            path: '/LucienShui',
            hovercard_url: '/users/LucienShui/hovercard',
          },
        },
        {
          total: 4,
          author: {
            id: 15308462,
            login: 'yangchuansheng',
            avatar: 'https://avatars2.githubusercontent.com/u/15308462?s=60&v=4',
            path: '/yangchuansheng',
            hovercard_url: '/users/yangchuansheng/hovercard',
          },
        },
        {
          total: 6,
          author: {
            id: 8605565,
            login: 'ysicing',
            avatar: 'https://avatars2.githubusercontent.com/u/8605565?s=60&v=4',
            path: '/ysicing',
            hovercard_url: '/users/ysicing/hovercard',
          },
        },
        {
          total: 6,
          author: {
            id: 19948780,
            login: 'PatHoo',
            avatar: 'https://avatars0.githubusercontent.com/u/19948780?s=60&v=4',
            path: '/PatHoo',
            hovercard_url: '/users/PatHoo/hovercard',
          },
        },
        {
          total: 8,
          author: {
            id: 33449123,
            login: 'jinnzy',
            avatar: 'https://avatars1.githubusercontent.com/u/33449123?s=60&v=4',
            path: '/jinnzy',
            hovercard_url: '/users/jinnzy/hovercard',
          },
        },
        {
          total: 29,
          author: {
            id: 38313565,
            login: 'oldthreefeng',
            avatar: 'https://avatars3.githubusercontent.com/u/38313565?s=60&v=4',
            path: '/oldthreefeng',
            hovercard_url: '/users/oldthreefeng/hovercard',
          },
        },
        {
          total: 194,
          author: {
            id: 4150905,
            login: 'cuisongliu',
            avatar: 'https://avatars0.githubusercontent.com/u/4150905?s=60&v=4',
            path: '/cuisongliu',
            hovercard_url: '/users/cuisongliu/hovercard',
          },
        },
        {
          total: 320,
          author: {
            id: 8912557,
            login: 'fanux',
            avatar: 'https://avatars1.githubusercontent.com/u/8912557?s=60&v=4',
            path: '/fanux',
            hovercard_url: '/users/fanux/hovercard',
          },
        },
      ]
      this.developer = developer.reverse()
    },
  },
}
</script> 