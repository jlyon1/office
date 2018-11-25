<template>
  <div>
    <section class="section">
      <section class="container">
        <div>
          <h1 class="title">
            Office Quick Reference
          </h1>
          <h2 class="subtitle">
            Find your favorite lines from The Office
          </h2>
          <div class="field has-addons ">
            <div class="control is-expanded">
              <input 
                v-model="text"
                class="input is-large" 
                type="text"
                @keyup.enter="grabData(text)">
            </div>
            <div class="control">
              <button 
                class="button is-large is-info"
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
    <div class="footer has-text-centered">
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
  methods: {
    grabData(query) {
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
.container {
  min-height: 100vh;
  display: flex;
  justify-content: center;
  text-align: center;
}

.title {
  font-family: 'Quicksand', 'Source Sans Pro', -apple-system, BlinkMacSystemFont,
    'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  display: block;
  font-weight: 300;
  font-size: 100px;
  color: #35495e;
  letter-spacing: 1px;
}

.subtitle {
  font-weight: 300;
  font-size: 42px;
  color: #526488;
  word-spacing: 5px;
  padding-bottom: 15px;
}

.links {
  padding-top: 15px;
}
</style>
