<template>
  <div class="home">
    <div class="phone-viewport">
      <md-toolbar>
        <md-button class="md-icon-button">
          <md-icon>menu</md-icon>
        </md-button>

        <h2 class="md-title" style="flex: 1">updoots</h2>

        <md-button class="md-icon-button">
          <md-icon>favorite</md-icon>
        </md-button>
      </md-toolbar>
      
      <div class="posts" >
        <md-list>
          <md-list-item v-for="post in posts" v-bind:key="post.id">
            <div class="md-list-text-container">
              {{post.message}}
              
            </div>
            <span>{{post.updoots}}</span>
              <md-button class="md-icon-button md-list-action" @click="updoot(post.id)">
                <md-icon>keyboard_arrow_up</md-icon>
              </md-button>
              
              <md-button class="md-icon-button md-list-action" @click="downdoot(post.id)">
                <md-icon>keyboard_arrow_down</md-icon>
              </md-button>
              
            <md-divider class="md-inset"></md-divider>
          </md-list-item>
        </md-list>
      </div>
    </div>
    <!-- this holds our dialog box for forcing auth -->
    <md-dialog-prompt
      :md-title="prompt.title"
      :md-ok-text="prompt.ok"
      :md-cancel-text="prompt.cancel"
      @open="onOpen()"
      @close="onClose()"
      v-model="prompt.value"
      ref="phoneauth">
    </md-dialog-prompt>
  </div>
</template>

<script>
// Import our API
import API from './api.js'

export default {
  name: 'home',
  data () {
    return {
      posts: [],
      prompt: {
        title: 'Enter your phone number:',
        ok: 'Ok!',
        cancel: 'Naw',
        value: '4128675309'
      }
    }
  },
  methods: {
    getMessages: function () {
      this.$http.get(API + '/message').then(response => {
        this.posts = response.body
      }, response => {
        // We need to get a token first
        // pretend we got the response
        this.posts = [{'id': 4, 'message': 'dae go to Ascend?', 'updoots': 2},
                        { 'id': 2, 'message': 'lmao', 'updoots': 1 },
                        { 'id': 3, 'message': 'this is a super cool app', 'updoots': 0 },
                        { 'id': 1, 'message': 'reply chug', 'updoots': 0 }]
        // this.$refs['phoneauth'].open()
      })
    },
    onClose: function (type) {
      this.$http.post(API + '/token', { phone: this.prompt.value }).then(response => {
        // We should have gotten a token
        this.getMessages()
      })
    },
    onOpen: function () {
      console.log('yay')
    },
    updoot: function (id) {
      this.posts.forEach(function (element) {
        if (element.id === id) {
          element.updoots += 1
        }
      }, this)
        /*
      this.$http.get(API + '/message/' + id + '/updoot').then(response => {
        this.posts.forEach(function(element) {
          if(element.id == id) {
            element.updoots += 1
          }
        }, this)
      }) */
      this.reorder()
    },
    downdoot: function (id) {
      this.posts.forEach(function (element) {
        if (element.id === id) {
          element.updoots -= 1
        }
      }, this)
      this.reorder()
    },
    reorder: function () {
      var temp = this.posts.slice(0)
      temp.sort(function (a, b) {
        return b.updoots - a.updoots
      })
      this.posts = temp
    }
  },
  created: function () {
    this.getMessages()
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
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
.posts {
  padding-left: 20px;
  padding-right: 20px;
}
.post {
  
}
</style>
