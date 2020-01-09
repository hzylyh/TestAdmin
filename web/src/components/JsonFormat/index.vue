<template>
  <div class="json-content">
    <div v-html="formatJson(jsonCont)">
    </div>
  </div>
</template>

<script>
export default {
  name: 'index',
  props: ['jsonCont', 'col'],
  data () {
    return {
    }
  },
  mounted () {
    this.replaceColVal()
  },
  methods: {
    replaceColVal () {
      let eachObj = this.jsonCont
      let eachCol = this.col.split('.')
      for (let i = 0; i < eachCol.length; i++) {
        if (i !== eachCol.length - 1) {
          eachObj = eachObj[eachCol[i]]
        } else {
          eachObj[eachCol[i]] = '<span style="color: red">c</span>'
        }
      }
    },
    formatJson (msg) {
      let rep = '~'
      let jsonStr = JSON.stringify(msg, null, rep).replace(/\\\"/g, '"')
      console.log(jsonStr)
      let str = ''
      for (let i = 0; i < jsonStr.length; i++) {
        let text2 = jsonStr.charAt(i)
        if (i > 1) {
          let text = jsonStr.charAt(i - 1)
          if (rep !== text && rep === text2) {
            str += '<br/>'
          }
        }
        str += text2
      }
      jsonStr = ''
      for (let i = 0; i < str.length; i++) {
        let text = str.charAt(i)
        if (rep === text) { jsonStr += '&nbsp;&nbsp;&nbsp;&nbsp;' } else {
          jsonStr += text
        }
        if (i === str.length - 2) { jsonStr += '<br/>' }
      }
      console.log(jsonStr)
      return jsonStr
    }
  }
}
</script>

<style lang="scss" scoped>
.json-content {
  text-align: left;
  .error-res {
    color: red;
  }
}
</style>
