<template>
  <div>
    <section class="section fullheight">
      <section class="container has-text-centered fullheight">
        <div>
          <h1 class="has-text-weight-bold is-size-1 my-font">
            Office Quick Reference
          </h1>
          <h2 class="subtitle is-size-3 my-font">
            Find your favorite lines from The Office
          </h2>
          <div class="field has-addons ">
            <div class="control is-expanded">
              <input 
                v-model="text"
                class="input" 
                type="text"
                @keyup.enter="grabData(text)">
            </div>
            <div class="control">
              <button 
                class="button is-info"
                @click="grabData(text)" >
                Search
              </button>
            </div>

          </div>
          <div
            v-if="results.length > 0"
            class="box has-text-left">
            <ul>
              <li 
                v-for="(res, idx) in sortedArray" 
                :key="idx">
                <p 
                  v-for="(line, idx) in res.Lines"
                  :key="idx">
                  - <strong>Season {{ res.Episode.Season }} Episode {{ res.Episode.Number }}:</strong>{{ line }}</p>
              </li>
            </ul>
          </div>
          <p v-if="results.length == 0">No results</p>
        </div>
      </section>
    </section>
    <div 
      style="bottom: 0; left: 0; right: 0; position: relative;" 
      class="footer has-text-centered has-background-white">
      <p>Made by <a href="https://jlyon.org">Joseph Lyon.</a> Check it out on <a href="https://github.com/jlyon1/office">Github</a></p>
    </div>
  </div>
</template>

<script>
import Logo from '~/components/Logo.vue'

export default {
  data() {
    return {
      results: [],
      text: ''
    }
  },
  computed: {
    sortedArray: function() {
      function compare(a, b) {
        return b.Rank - a.Rank
      }
      return this.results.slice(0).sort(compare)
    }
  },
  mounted() {
    if (this.$router.app._route.query.search != undefined) {
      this.text = this.$router.app._route.query.search
      this.grabData(this.text)
    }
  },
  methods: {
    grabData(query) {
      this.$router.push('/?search=' + encodeURIComponent(query))

      fetch(
        'http://localhost:8080/office/search?query=' + encodeURIComponent(query)
      )
        .then(resp => {
          return resp.json()
        })
        .then(val => {
          if (val === null) {
            this.results = []
            return
          }
          this.results = val
        })
    }
  }
}
</script>
<style>
.my-font {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
}

.links {
  padding-top: 15px;
}

.fullheight {
  min-height: 100%;
}
</style>
