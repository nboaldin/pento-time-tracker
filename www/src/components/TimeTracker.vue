<template>
  <v-container>
    <v-row>
      <v-col cols="3">
        <Form @userFormSubRes="updateUserData" />
        <SessionForm
          @sessionFormSubRes="updateAfterSessionStart"
          :user="this.id"
        />
      </v-col>
      <v-col>
        <v-sheet min-height="70vh" rounded="lg">
          <v-list>
            <h4>{{ firstName + " " + lastName }}</h4>

            <v-list-item v-for="session in sessions" :key="session.id">
              <v-list-item-content>
                <v-list-item-title v-text="session.name"></v-list-item-title>
                <v-list-item-subtitle
                  v-text="session.start"
                ></v-list-item-subtitle>
                <v-list-item-subtitle
                  v-text="session.end"
                ></v-list-item-subtitle>
              </v-list-item-content>
              <v-list-item-action>
                <v-btn
                  v-show="sessionNotStopped(session.end)"
                  @click="stopSession(session.id)"
                >
                  Stop
                  <v-icon color="red lighten-1">mdi-stop</v-icon>
                </v-btn>
              </v-list-item-action>
            </v-list-item>
          </v-list>
        </v-sheet>
      </v-col>
    </v-row>
    <Overview :sess="this.sessions" />
  </v-container>
</template>

<script>
import Form from "./Form";
import SessionForm from "./SessionForm";
import Overview from "./Overview";

export default {
  components: {
    Form,
    SessionForm,
    Overview
  },

  data: () => ({
    id: "",
    firstName: "",
    lastName: "",
    sessions: []
  }),
  methods: {
    updateUserData(data) {
      this.id = data.id;
      this.firstName = data.firstName;
      this.lastName = data.lastName;
      this.sessions = data.timeTrackingSessions;
    },
    updateAfterSessionStart(data) {
      this.id = data.userId;
      this.firstName = data.firstName;
      this.lastName = data.lastName;
      this.sessions = data.sessions;
    },
    stopSession(uuid) {
      console.log(`Stop clicked: ${uuid}`);
      this.axios
        .post(
          `/user/${this.id}/session_stop`,
          JSON.stringify({
            id: uuid
          })
        )
        .then(r => {
          console.log(r.data);
          const { timeTrackingSessions } = r.data;
          this.sessions = timeTrackingSessions;
        });
    },
    sessionNotStopped(endDate) {
      return endDate === "0001-01-01T00:00:00Z";
    }
  }
};
</script>
