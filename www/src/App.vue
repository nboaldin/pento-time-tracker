<template>
  <v-app>
    <v-main class="grey lighten-3">
      <TimeTracker />
    </v-main>
  </v-app>
</template>

<script>
import TimeTracker from "./components/TimeTracker";

export default {
  name: "App",

  components: {
    TimeTracker
  },

  data: () => ({
    //
  })
};
</script>
