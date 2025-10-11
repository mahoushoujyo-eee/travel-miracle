<script setup>
import { onMounted, onUnmounted } from "vue";
import AMapLoader from "@amap/amap-jsapi-loader";

let map = null;

onMounted(() => 
{
  window._AMapSecurityConfig = {
    securityJsCode: "bef6bdac946bcf00ce8781b63e50affc",
  };

  AMapLoader.load(
  {
    key: "0efec05f109e098f43372e271e03a3b6", // 申请好的Web端开发者Key，首次调用 load 时必填
    version: "2.0", // 指定要加载的 JSAPI 的版本，缺省时默认为 1.4.15
    plugins: ["AMap.ToolBar", "AMap.Scale", "AMap.PlaceSearch", "AMap.Geolocation", "AMap.PlaceSearch"], // 需要使用的的插件列表，要先在此处声明
  })
  .then((AMap) => 
  {
    map = new AMap.Map("container", 
    {
        // 设置地图容器id
        viewMode: "2D", // 是否为3D地图模式
        zoom: 14, // 初始化地图级别
        center: [116.397428, 39.90923], // 初始化地图中心点位置
    });
    const toolbar = new AMap.ToolBar(); //创建工具条插件实例
    map.addControl(toolbar); //添加工具条插件到页面
    const scale = new AMap.Scale();
    map.addControl(scale);
    // var placeSearch = new AMap.PlaceSearch();
    // map.addControl(placeSearch); //添加地点搜索插件到页面

    // 添加定位插件 needAddress 为 true 时，会返回定位地址详细信息
    const geolocation = new AMap.Geolocation({"needAddress":true});
    geolocation.getCurrentPosition((status, result)=>{console.log(status, result)})

    // 地点搜索插件
    const PlaceSearchOptions = { //设置PlaceSearch属性
        city: "南京", //城市
        type: "", //数据类别
        pageSize: 10, //每页结果数,默认10
        pageIndex: 1, //请求页码，默认1
        extensions: "base" //返回信息详略，默认为base（基本信息）
    };
    const MSearch = new AMap.PlaceSearch(PlaceSearchOptions); //构造PlaceSearch类
    MSearch.search('南京师范大学', (status, result)=>{console.log(status, result)}); //关键字查询
  })
  .catch((e) => 
  {
      console.log(e);
  });
});

onUnmounted(() => 
{
  map?.destroy();
});
</script>

<template>
  <div id="container"></div>
</template>

<style scoped>
#container {
  width: 100%;
  height: 800px;
}
</style>
