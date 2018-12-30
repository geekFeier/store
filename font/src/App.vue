<template>
  <div id="app">
    <Menu mode="horizontal" :theme="theme1" active-name="1">
      <Row>
        <Col span="19">
        <MenuItem name="1">
          <a target="_blank" href="http://store.lameleg.com">
            <li>SealYun</li>
          </a>
        </MenuItem>
        <MenuItem name="2">
          <a target="_blank" href="https://www.sealyun.com">
           <li>Blog</li>
          </a>
        </MenuItem>
        <MenuItem name="3">
          <a target="_blank" href="http://www.sealyun.com/pro/products">
            <li>Docs</li>
          </a>
        </MenuItem>
        </Col>
        <Col span="5">
        <MenuItem name="4">
          <Tooltip content="分享收入可提现金额">
            <li>
              <i class="ivu-icon ivu-icon-logo-yen">
                <Badge :count="100">&nbsp;&nbsp;&nbsp;&nbsp;</Badge>
              </i>
            </li>
          </Tooltip>
        </MenuItem>
        <MenuItem name="5">
          <Tooltip content="请使用github账户登录">
            <a
              href="https://github.com/login/oauth/authorize?client_id=89c1b05d77fb1c92a1ef&scope=user:email"
              target="_blank"
              rel="noopener"
            >
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
    <HelloWorld msg="Welcome to SealYun"/>
  </div>
</template>

<script>
import HelloWorld from "./components/HelloWorld.vue";

export default {
  data () {
    return {
      theme1:'light'
    }
  },
  name: "app",
  components: {
    HelloWorld
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
</style>
