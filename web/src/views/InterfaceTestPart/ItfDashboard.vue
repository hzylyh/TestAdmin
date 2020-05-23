<!--
 * @Description: 接口测试部分Dashboard
 * @Author: John Holl
 * @Github: https://github.com/hzylyh
 -->
<template>
  <el-row>
    <panel-group :panel-data="panelData" @handleSetLineChartData="handleSetLineChartData" />
    <el-row style="background:#fff;padding:16px 16px 0;margin-bottom:32px;">
      <line-chart :chart-data="lineChartData" />
    </el-row>
  </el-row>
</template>

<script>
import PanelGroup from './components/PanelGroup'
import LineChart from './components/LineChart'
import { getCaseRunInfo } from '@/api/dashboard'

export default {
  name: 'ItfDashboard',
  components: {
    PanelGroup,
    LineChart
  },
  data () {
    return {
      panelData: {
        totalCase: 2000
      },
      caseRunInfo: null,
      lineChartData: {
        xData: [],
        expectedData: [100, 120, 161, 134, 105, 160, 165],
        actualData: [120, 82, 91, 154, 162, 140, 145]
      }
      // lineChartData: {
      //   newVisitis: {
      //     xData: [],
      //     expectedData: [100, 120, 161, 134, 105, 160, 165],
      //     actualData: [120, 82, 91, 154, 162, 140, 145]
      //   },
      //   messages: {
      //     xData: [],
      //     expectedData: [200, 192, 120, 144, 160, 130, 140],
      //     actualData: [180, 160, 151, 106, 145, 150, 130]
      //   },
      //   purchases: {
      //     xData: [],
      //     expectedData: [80, 100, 121, 104, 105, 90, 100],
      //     actualData: [120, 90, 100, 138, 142, 130, 130]
      //   },
      //   shoppings: {
      //     xData: [],
      //     expectedData: [130, 140, 141, 142, 145, 150, 160],
      //     actualData: [120, 82, 91, 154, 162, 140, 130]
      //   }
      // }
    }
  },
  mounted () {
    this.getCaseRunInfo()
  },
  methods: {
    getCaseRunInfo () {
      getCaseRunInfo().then(response => {
        console.log(response)
        this.lineChartData.actualData = response.success
        this.lineChartData.expectedData = response.fail
        this.lineChartData.xData = response.beginTime
      })
    },
    handleSetLineChartData (type) {
      // this.lineChartData = lineChartData[type]
      this.lineChartData = 'newVisitis'
    }
  }
}
</script>

<style scoped>

</style>
