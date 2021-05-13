<template>
  <div class="logviewer">
    <div v-for="(log,i) in logs" v-bind:key="i">
      <span class="status-point" style=" background-color:#67C23A" ></span>
      <span class="status-point" style=" background-color:#E6A23C" ></span>
      <span>{{ log.time }}</span>
      <span>{{ log.level }}</span>
      <span>{{ log.trace_id }}</span>
    </div>
  </div>
</template>

<script>
export default {
  name: 'LogViewer',
  data() {
    return {
      logs: []
    }
  },
  props: {
    msg: String
  }, mounted() {
    let app = this
    var loc = window.location;
    var uri = 'ws:';

    if (loc.protocol === 'https:') {
      uri = 'wss:';
    }
    uri += '//' + loc.host;
    uri += '/ws';

    let ws;
    ws = new WebSocket(uri)

    ws.onopen = function () {
      console.log('Connected')
    }

    ws.onmessage = function (evt) {

      app.logs.push(JSON.parse(evt.data) )

    }

    setInterval(function () {
      ws.send('Hello, Server!');
    }, 1000);

  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
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

.status-point {
  display: inline-block;
  width: 6px;
  height: 6px;
  border-radius: 50%;
}

</style>
