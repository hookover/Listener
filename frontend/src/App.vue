<template>
  <div id="app" v-wechat-title="title">
    <div class="search-div">
      <v-select class="search" v-model="selected" @search="search" :options="search_data"></v-select>
      <span class="book_number">当前书籍数量：{{ this.book_number }} 部</span>
    </div>

    <aplayer
      ref="aplayer"
      :audio="books"
      :autoplay="auto_play"
      :listMaxHeight="listMaxHeight"
      :LrcType="0"
    />

  </div>
</template>

<script>
/* eslint-disable */
  import Vue from 'vue';
  import APlayer from '@moefe/vue-aplayer';
  import axios from 'axios'
  // import Aplayer from 'vue-aplayer'
  import vSelect from 'vue-select'

  Vue.use(APlayer);

  export default {
    name: 'App',
    components: {
      vSelect
    },
    data: function () {
      return {
        base_url: 'http://123.206.56.200:8080/',
        title: '听书',
        volume: 1,
        muted: false,
        books: [],
        book_number: 0,
        search_data: [],
        selected: "",
        auto_play: true,
        listMaxHeight:500
      }
    },
    mounted:function() {
      this.$watch("$refs.aplayer.currentMusic", (n, o) => {
          this.title = n.name;
        }
      );

      this.getMp3List();

      if(this.$refs.aplayer.currentMusic) {
        this.title = this.$refs.aplayer.currentMusic.name
      }
    },
    watch: {
      selected(n, o) {
        if (n) {
          if (this.$refs.hasOwnProperty('aplayer')) {
            this.$refs.aplayer.switch(n.label);
            this.selected = "";
          }
        }
      }
    },
    computed: {},
    methods: {
      getRandomInt: function (n, m) {
        return Math.floor(Math.random() * (m - n + 1) + n);
      },
      getMp3List: function () {
        axios.get('/book/list').then(respose => {
          if (respose.data.code === 200) {
            for (let n of respose.data.data) {
              this.books.push({
                name: n.Name,
                artist: n.Name,
                url: this.base_url + n.Path,
              });
            }
            this.book_number = this.books.length;
          }
        })
      },
      search: function (search, loading) {
        this.search_data = [];
        for (let i = 0; i < this.book_number; i++) {
          if (this.books[i].name.indexOf(search) >= 0) {
            this.search_data.push({
              label: this.books[i].name,
            })
          }
        }
      },
    }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
  }

  #app .search-div {
    /*text-align: left;*/
    padding: 30px 5px 20px;
    display: inline-block;
    width: 100%;
  }

  #app .book_number {
    float: right;
    margin-right: 15px;
  }

  #app .search {
    min-width: 50%;
    /*max-width: 50%;*/
    float: left;
  }
</style>
