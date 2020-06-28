<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .layout{
        border: 0px solid #d7dde4;
        background: #f5f7f9;
        position: relative;
        border-radius: 4px;
        overflow: hidden;
    }
    .layout-header-bar{
        background: #fff;
        box-shadow: 0 1px 1px rgba(0,0,0,.1);
    }
</style>
<template>
    <div class="layout">
        <Sider :style="{position: 'fixed', height: '100vh', left: 0, overflow: 'auto'}">
            <!-- <router-link to="/home">
                <Icon type="ios-home"></Icon>
                Home
            </router-link> -->
            <Menu active-name="1-2" theme="dark" width="auto" :open-names="['1']">
                <MenuItem name="4-1" to="/home">
                    <Icon type="ios-home"></Icon>
                    Home
                </MenuItem>
                <Submenu name="1">
                    <template slot="title">
                        <Icon type="ios-keypad"></Icon>
                        Item
                    </template>
                    <MenuItem name="1-1">Anime</MenuItem>
                    <MenuItem name="1-2">Manga</MenuItem>
                </Submenu>
                <Submenu name="2">
                    <template slot="title">
                        <Icon type="md-cloud-download" />
                        Download
                    </template>
                    <MenuItem name="2-1" to="/video_download">V</MenuItem>
                    <MenuItem name="2-2" to="/thunder_download">T</MenuItem>
                </Submenu>
                <Submenu name="3">
                    <template slot="title">
                        <Icon type="md-heart-outline" />
                        Other
                    </template>
                    <MenuItem name="3-1" to="/hello">Hello</MenuItem>
                </Submenu>
            </Menu>
        </Sider>
        <Layout :style="{marginLeft: '200px'}">
            <Header :style="{background: '#fff', boxShadow: '0 2px 3px 2px rgba(0,0,0,.1)'}"></Header>
            <Content :style="{padding: '0 16px 16px'}">
                <Breadcrumb :style="{margin: '16px 0'}">
                    <BreadcrumbItem>Home</BreadcrumbItem>
                    <BreadcrumbItem>Components</BreadcrumbItem>
                    <BreadcrumbItem>Layout</BreadcrumbItem>
                </Breadcrumb>
                <Card>
                    <!-- 路由出口 -->
                    <!-- 路由匹配到的组件将渲染在这里 -->
                    <router-view></router-view>
                </Card>
            </Content>
        </Layout>
    </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      title: 'JiHan'
    }
  }
}
</script>
