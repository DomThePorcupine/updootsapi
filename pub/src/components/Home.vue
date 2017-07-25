<template>
  <div class="home">
    <div class="phone-viewport">
      <md-whiteframe md-tag="md-toolbar" md-elevation="2" md-theme="light-blue" class="md-small">
        <div class="md-toolbar-container">
          <md-button class="md-icon-button" @click="$refs.sidenav.toggle()">
            <md-icon>menu</md-icon>
          </md-button>
          <h2>updoots</h2>
          <span style="flex: 1"></span>
          
          <md-button @click="$refs['auth'].open()" class="md-icon-button">
          <md-icon>account_box</md-icon>
        </md-button>
        <md-button @click="$refs['register'].open()" class="md-icon-button">
          <md-icon>person_add</md-icon>
        </md-button>
        </div>
      </md-whiteframe>

     
      
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
      ref="auth">
    </md-dialog-prompt>

    <md-dialog-prompt
      :md-title="register.title"
      :md-ok-text="register.ok"
      :md-cancel-text="register.cancel"
      @open="onRegOpen()"
      @close="onRegClose()"
      v-model="register.value"
      ref="register">
    </md-dialog-prompt>

    <md-dialog-prompt
      :md-title="createpost.title"
      :md-ok-text="createpost.ok"
      :md-cancel-text="createpost.cancel"
      @open="onCreateOpen()"
      @close="onCreateClose()"
      v-model="createpost.value"
      ref="createpost">
    </md-dialog-prompt>

    <md-snackbar :md-position="'bottom center'" ref="snackbar" :md-duration="4000">
      <span>You have succefully registered.</span>
      <md-button class="md-accent" md-theme="light-blue" @click="$refs.snackbar.close()">Close</md-button>
    </md-snackbar>

    
      <md-button @click="$refs['createpost'].open()" class="md-fab dom-float md-fab-bottom-right">
        <md-icon>add</md-icon>
      </md-button>
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
        title: 'Enter your userid:',
        ok: 'Ok!',
        cancel: 'Naw',
        value: 'user'
      },
      register: {
        title: 'Enter a new userid',
        ok: 'Ok!',
        cancel: 'Naw',
        value: 'user@pitt.edu'
      },
      createpost: {
        title: 'Enter your post',
        ok: 'Post!',
        cancel: 'Cancel',
        value: ''
      }
    }
  },
  methods: {
    getMessages: function () {
      this.$http.get(API + '/message').then(response => {
        this.posts = JSON.parse(response.body)
      }, response => {
        this.$refs['auth'].open()
      })
    },
    onClose: function () {
      this.$http.post(API + '/token', JSON.stringify({ userid: this.prompt.value }), {
        headers: {
          'Content-Type': 'text/plain',
          'Accept': 'application/json'
        },
        options: {
          'withCredentials': true
        }
      }).then(response => {
        // We should have gotten a token
        this.getMessages()
      })
    },
    onOpen: function () {
      console.log('yay')
    },
    onRegOpen: function () {
      console.log('yay')
    },
    onRegClose: function () {
      this.$http.post(API + '/register', { userid: this.register.value }).then(response => {
        this.$refs.snackbar.open()
      })
    },
    onCreateOpen: function () {
      console.log('yay')
    },
    onCreateClose: function () {
      if (this.createpost.value !== '') {
        this.$http.post(API + '/message', { message: this.createpost.value.toString('utf8') }).then(response => {
          this.getMessages()
          this.createpost.value = ''
        })
      }
    },
    updoot: function (id) {
      this.$http.post(API + '/doot', { doot: 1, message: id }).then(response => {
        // Here
        this.getMessages()
      })
    },
    downdoot: function (id) {
      this.$http.post(API + '/doot', { doot: 0, message: id }).then(response => {
        // Here
        this.getMessages()
      })
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
.complete-example {
  height: 540px;
  display: flex;
  flex-flow: column;
  position: relative;
  overflow: hidden;
  z-index: 1;

  .md-fab {
    margin: 0;
    position: absolute;
    bottom: -20px;
    left: 16px;
  }
  .dom-float {
    position:fixed;
    right:0;
    bottom:0;
  }

  .md-title {
    color: #fff;
  }

  .md-list {
    overflow: auto;
  }

  .md-list-action .md-icon {
    color: rgba(#000, .26);
  }

  .md-avatar-icon .md-icon {
    color: #fff !important;
  }

  .md-sidenav .md-list-text-container > :nth-child(2) {
    color: rgba(#fff, .54);
  }

  .md-account-header {
    .md-list-item:hover .md-button:hover {
      background-color: inherit;
    }

    .md-avatar-list .md-list-item-container:hover {
      background: none !important;
    }
  }
}
</style>
