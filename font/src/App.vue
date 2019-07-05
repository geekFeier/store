<template>
  <div id="app">
    <Menu mode="horizontal" :theme="theme1" active-name="1">
      <Row>
        <Col span="21">
          <MenuItem name="1">
            <a href="http://store.lameleg.com">
              <li>SealYun</li>
            </a>
          </MenuItem>
          <MenuItem name="2">
            <a target="_blank" href="https://sealyun.com">
              <li>博客</li>
            </a>
          </MenuItem>
          <MenuItem name="3">
            <a target="_blank" href="https://sealyun.com/post/docs/">
              <li>文档</li>
            </a>
          </MenuItem>
          <MenuItem name="3">
            <a target="_blank" href="https://sealyun.com/post/referrer/">
              <li>加入营销</li>
            </a>
          </MenuItem>
          <MenuItem name="3">
            <a target="_blank" href="https://mp.weixin.qq.com/s/Ra722VtdLitDbM0GExom6A">
              <li>售后支持</li>
            </a>
          </MenuItem>
          <MenuItem name="3">
            <a target="_blank" href="https://shop929htt07.market.aliyun.com/page/productlist.html?cId=53366009">
              <li>阿里云市场</li>
            </a>
          </MenuItem>
        <Submenu name="4">
            <template slot="title">
                <Icon type="logo-github" />
                开源项目
            </template>
            <MenuGroup title="使用">
              <MenuItem name="0">
                <a target="_blank" href="https://github.com/fanux/sealos">
                  <li>kubernetes一键HA</li>
                </a>
              </MenuItem>
              <MenuItem name="1">
                <a target="_blank" href="https://github.com/fanux/fist">
                  <li>轻量级kubernetes管理工具</li>
                </a>
              </MenuItem>
              <MenuItem name="1">
                <a target="_blank" href="https://github.com/fanux/lhttp">
                  <li>好用的websocket框架</li>
                </a>
              </MenuItem>
            </MenuGroup>
        </Submenu>

        </Col>
        <Col span="3">
          <MenuItem name="4">
            <Tooltip :content="'分享收入可提现金额,您可提现'+amount+'元'">
              <li @click="payeeFormCheck">
                <Modal v-model="payeeForm" @on-ok="ok" title="设置收款支付宝账号与提现密码" @on-cancel="cancel">
                  <div class="payee">
                    <label>提现账号</label>
                    <input v-model="account" placeholder="收款支付宝账号">
                    <br>
                    <label>提现密码</label>
                    <input v-model="passwd" placeholder="提现安全密码">
                    <br>
                    <label>确认密码</label>
                    <input v-model="passwdCheck" placeholder="切勿使用支付宝密码">
                    <br>
                  </div>
                </Modal>
                <i class="ivu-icon ivu-icon-logo-yen">
                  <Badge :count="amount">&nbsp;&nbsp;&nbsp;&nbsp;</Badge>
                </i>
              </li>
            </Tooltip>
          </MenuItem>

          <MenuItem name="5">
            <!--img v-if="avata" :src="avata_url" style="border-radius:50%;width:30px;height:30px;cursor:pointer;margin-top:12px;"></img-->
            <Avatar v-if="avata" :src="avata_url"/>
            <Tooltip v-else content="请使用github账户登录">
              <!-- useless, please redirect on backend-->
              <a :href="loginurl" target="_blank" rel="noopener">
                <li>
                  <i class="ivu-icon ivu-icon-logo-github"></i>
                  Login
                </li>
              </a>
            </Tooltip>
          </MenuItem>
        </Col>
      </Row>
    </Menu>
    <Row>
      <Col span="8">
        <h1 class="payedUser">已购用户</h1>
        <label v-for="u in payedUser">
          <a :href="'https://github.com/'+u.login" target="_blank">
            <Avatar class="payedUser" :src="u.avatar_url"/>
          </a>
        </label>
      </Col>
      <Col span="8">
        <HelloWorld msg="Welcome to SealYun"/>
      </Col>
      <Col span="8">
        <div style="text-align:left;" id="gitalk-container"></div>
      </Col>
    </Row>
    <!--router-view></router-view-->
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";
import VueCookies from "vue-cookies";

import "gitalk/dist/gitalk.css";
import Gitalk from "gitalk";

/*
const gitalk = new Gitalk({
  clientID: '98478b0f6bfacff7cdf0',
  clientSecret: 'a882c5ed737d7453392b83c6b25e232a9d859d03',
  repo: 'https://github.com/fanux/store',
  owner: 'fanux',
  admin: ['fanux'],
  id: "gitalk.store.lameleg.com",      // Ensure uniqueness and length less than 50
  distractionFreeMode: false  // Facebook-like distraction free mode
})

gitalk.render('gitalk-container')
*/

export default {
  mounted() {
    const gitalk = new Gitalk({
      clientID: "98478b0f6bfacff7cdf0",
      clientSecret: "a882c5ed737d7453392b83c6b25e232a9d859d03",
      repo: "gitalk",
      owner: "fanux",
      admin: ["fanux"],
      id: window.location.pathname, // Ensure uniqueness and length less than 50
      distractionFreeMode: false // Facebook-like distraction free mode
    });

    gitalk.render("gitalk-container");
  },

  data() {
    var a = {
      account: "",
      passwd: "",
      passwdCheck: "",

      amount: 1,
      theme1: "light",
      avata: false,
      avata_url: "",
      payeeForm: false,

      payedUser: [],

      loginurl:
        "https://github.com/login/oauth/authorize?client_id=89c1b05d77fb1c92a1ef&scope=user:email"
    };

    this.$http
      .get(
        "http://store.lameleg.com:8080/loginless/pro/kubernetes1.13.1/payed",
        {
          credentials: true
        }
      )
      .then(
        function(res) {
          for (var i = 0; i < res.data.length; i++) {
            if (res.data[i] == null) {
              continue;
            }
            for (var j = i + 1; j < res.data.length; j++) {
              if (res.data[j] == null) {
                continue;
              }
              if (res.data[i].login == res.data[j].login) {
                // res.data.splice(j - 1, 1);
                res.data[j].avata_url =
                  "https://avatars2.githubusercontent.com/u/8912557?v=4";
              }
            }
          }

          a.payedUser = res.data;
        },
        function(res) {}
      );

    if (typeof this.$route.query.referrer != "undefined") {
      // a.loginurl += "&redirect_uri=http://store.lameleg.com/referrer/" + this.$route.query.referrer;
      VueCookies.set("referrer", this.$route.query.referrer);
    }
    if (VueCookies.get("referrer") != null) {
      console.log("cookie", VueCookies.get("referrer"));
    }
    this.$http
      .get("http://store.lameleg.com:8080/loginless/user/payee", {
        credentials: true
      })
      .then(
        function(res) {
          a.amount = res.data.Amount;
          if (typeof a.amount == "undefined") {
            a.amount = 0;
          }
          console.log(res.data);
        },
        function(res) {
          console.log(res.data);
        }
      );

    this.$http
      .get("http://store.lameleg.com:8080/loginless/info/user", {
        credentials: true
      })
      .then(
        function(res) {
          a.avata_url = res.data.avatar_url;
          //              a.avata_url="https://avatars2.githubusercontent.com/u/8912557?v=4"
          if (typeof a.avata_url != "undefined") {
            a.avata = true;
          }
          console.log(res.data, "avata:", a);
        },
        function(res) {
          console.log(res.data);
        }
      );

    return a;
  },
  name: "app",
  components: {
    HelloWorld
  },
  methods: {
    ok() {
      if (this.passwd != this.passwdCheck) {
        this.$Message.info("两次密码设置不相同");
        return;
      }

      var para = {
        PayeeAccount: this.account,
        Passwd: this.passwd
      };

      this.$http
        .put("http://store.lameleg.com:8080/user/info/payee", para, {
          credentials: true,
          responseType: "json"
        })
        .then(
          function(res) {
            console.log(res);
            this.$Message.info(res.data.Reason);
          },
          function(res) {
            this.$Message.info(res.data.Reason);
          }
        );
    },
    cancel() {
      this.$Message.info("取消设置");
    },
    payeeFormCheck: function(event) {
      this.$http
        .post(
          "http://store.lameleg.com:8080/user/info/withdraw",
          {},
          { credentials: true }
        )
        .then(
          function(res) {
            this.payeeForm = true;
            if (res.data.Amount == 0) {
              console.log("withdraw is 0");
            }
            if (res.data.PayeeAccount == "" || res.data.Passwd == "") {
              // if not set, else withdraw
              this.payeeForm = true;
            }
            this.account = res.data.PayeeAccount;
            console.log(res.data);
          },
          function(res) {
            console.log("user info withdraw failed", res.data);
            this.payeeForm = true;
          }
        );
    }
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
.payee {
  margin: 10px;
}
.payee label {
  margin: 10px;
}
.payee input {
  margin: 10px;
}
.payedUser {
  margin: 10px;
}
</style>
