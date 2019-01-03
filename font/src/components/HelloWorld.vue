<template>
  <div class="hello">
    <h1>Welcom to SealYun</h1>
    <p>kubernetes生态应用市场
      <br>容器可复用软件销售平台
      <br>kubernetes集群离线安装包，一键安装，HA安装
      <br>
      <!--a href="https://github.com/login/oauth/authorize?client_id=89c1b05d77fb1c92a1ef&scope=user:email" target="_blank" rel="noopener">login github</a-->
    </p>
    <h3>商品列表</h3>
    <div id="pro-link">
      <ul>
        <li>
          <a :href="time" target="_blank" rel="noopener">kubernetes1.13.1 (元旦一折优惠中 5元)</a>
          <Tooltip content="获取专有分享链接，享受60%交易提成">
            <Button @click="share = true" id="sharelink" type="success">推广链接</Button>
            <Modal v-model="share" title="专有分享链接 - 通过sealyun赚钱" @on-ok="ok" @on-cancel="cancel">
              <p>{{ shareLink }}</p>
              <p>任何用户通过上面链接访问网站并成功交易您将获得交易的60%提成</p>
              <p>如嵌入自己的markdown文档中 [kubernetes离线安装仅需三步]({{ shareLink }})</p>
              <p>功能研发中</p>
            </Modal>
          </Tooltip>
        </li>
      </ul>
    </div>
    <h3>LINKS</h3>
    <ul>
      <li>
        <img src="https://sealyun.com/kubernetes-qrcode.jpg">
      </li>
      <br>公众号
      <br>
      <li>联系方式：fhtjob@hotmail.com</li>
    </ul>
  </div>
</template>

<script  type="text/javascript">
export default {
  name: "HelloWorld",
  props: {
    msg: String
  },
  methods: {
    ok() {
      this.$Message.info("Clicked ok");
    },
    cancel() {
      this.$Message.info("Clicked cancel");
    }
  },
  data: function() {
    var d = {
      time:
        "http://store.lameleg.com:8080/pro/kubernetes1.13.1?time=" +
        new Date().getTime()+"&referrer=",
      shareLink: "",
      share: false,
    };

    this.$http
      .get("http://store.lameleg.com:8080/loginless/info/user", {
        credentials: true
      })
      .then(
        function(res) {
          d.shareLink = "http://store.lameleg.com?referrer="+res.data.login;
          console.log(res.data, "info:", d);
        },
        function(res) {
          console.log(res.data);
        }
      );

    return d;
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
#sharelink {
  margin: 10px;
}
</style>
