<template>
  <div class="home">
    <div class="phone-viewport">
      <md-toolbar>
        <md-button class="md-icon-button">
          <md-icon>menu</md-icon>
        </md-button>

        <h2 class="md-title" style="flex: 1">Default</h2>

        <md-button class="md-icon-button">
          <md-icon>favorite</md-icon>
        </md-button>
      </md-toolbar>
      {{messages}}
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
      messages: [],
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
        this.messages = response.body
      }, response => {
        // We need to get a token first
        this.$refs['phoneauth'].open()
      })
    },
    deleteBatch: function (id) {
      this.$http.delete(API + '/api/batch/' + id).then(function (response) {
        this.getBatches()
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
</style>
