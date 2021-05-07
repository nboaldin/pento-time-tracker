<template>
  <v-container>
    <v-row>
      <v-col v-for="category in categories" :key="category.heading">
        <v-list>
          <h4>{{ category.heading }}</h4>
          <v-list-item v-for="session in startingSessions" :key="session.id">
            <v-list-item-content>
              <v-list-item-title v-text="session.name"></v-list-item-title>
              <v-list-item-subtitle
                v-text="session.start"
              ></v-list-item-subtitle>
              <v-list-item-subtitle v-text="session.end"></v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list>
      </v-col>
      <v-col></v-col>
      <v-col></v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  props: {
    sess: Array
  },

  data: () => ({
    userId: "",
    startingSessions: [],
    categories: {
      today: {
        date: "",
        heading: "Today",
        sessions: []
      },
      thisWeek: {
        startDate: "",
        endDate: "",
        heading: "This Week",
        sessions: []
      },
      thisMonth: {
        startDate: "",
        endDate: "",
        heading: "This Month",
        sessions: []
      }
    }
  }),

  watch: {
    sess: {
      deep: true,
      handler() {
        this.startingSessions = [];

        // this.sess.forEach(e => {
        //   e.start = new Date(e.start);
        //   e.end = new Date(e.end);
        //   this.startingSessions.push(e);
        // });
        this.startingSessions.push(...this.sess);
        console.log(
          "sessions in overview after watch and push",
          this.startingSessions
        );
      }
    }
  },

  methods: {
    todaySessions() {}
  }
};
</script>

<style scoped></style>
