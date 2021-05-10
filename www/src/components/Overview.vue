<template>
  <v-container>
    <v-row>
      <v-col v-for="category in categories" :key="category.heading">
        <v-list>
          <h4>{{ category.heading }}</h4>
          <v-list-item v-for="session in category.sessions" :key="session.id">
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
      handler: function() {
        // this.startingSessions = [];

        this.startingSessions = this.sess;
        this.todaySessions();
        this.thisMonthSessions();
        this.thisWeekSessions();

        // this.sess.forEach(e => {
        //   e.start = new Date(e.start);
        //   e.end = new Date(e.end);
        //   this.startingSessions.push(e);
        // });
        // this.startingSessions.push(...this.sess);
        console.log(
          "sessions in overview after watch and push",
          this.startingSessions
        );
      }
    }
  },

  methods: {
    todaySessions() {
      const now = Date.now();
      const today = new Date(now);
      const todayDate = today.getDate();
      const todayMonth = today.getMonth();
      const todayYear = today.getFullYear();

      const todayArr = this.startingSessions.filter(function(sess) {
        const sessionDate = new Date(sess.start);

        if (
          sessionDate.getDate() === todayDate &&
          sessionDate.getMonth() === todayMonth &&
          sessionDate.getFullYear() === todayYear
        ) {
          return true;
        }
        return false;
      });
      this.categories.today.sessions = todayArr;
    },
    thisWeekSessions() {
      // This is incomplete and untested. I could have copied someones getWeek prototype but I wanted to think through it.
      const now = Date.now();
      const today = new Date(now);
      console.log("Today:", today);
      const todayDateDay = today.getDay();
      const thisWeek = [];

      //This will cause a duplicate in the array
      thisWeek.push(new Date(Date.now()));

      for (let i = todayDateDay; i >= 1; i--) {
        console.log("Subtract days ran:", i);
        thisWeek.push(new Date(today.setDate(today.getDate() - 1)));
      }
      for (let i = todayDateDay; i <= 5; i++) {
        console.log("Add days ran:", i);
        thisWeek.push(new Date(today.setDate(today.getDate() + 1)));
      }

      console.log("This week array:", thisWeek);
      //Still need to check if each session matches the thisWeek array and display it.
      // this.categories.thisWeek.sessions = thisWeekArr;
    },
    thisMonthSessions() {
      const now = Date.now();
      const today = new Date(now);
      const todayMonth = today.getMonth();
      const todayYear = today.getFullYear();

      const thisMonthArr = this.startingSessions.filter(function(sess) {
        const sessionDate = new Date(sess.start);

        if (
          sessionDate.getMonth() === todayMonth &&
          sessionDate.getFullYear() === todayYear
        ) {
          return true;
        }
        return false;
      });
      this.categories.thisMonth.sessions = thisMonthArr;
    }
  }
};
</script>

<style scoped></style>
