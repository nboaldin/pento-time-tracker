<template>
  <form>
    <v-text-field
      v-model="sessionName"
      :error-messages="sessionNameErrors"
      :counter="20"
      label="Session Name"
      required
      @input="$v.sessionName.$touch()"
      @blur="$v.sessionName.$touch()"
    ></v-text-field>
    <v-btn class="mr-4" @click="startSession">
      start
    </v-btn>
  </form>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, maxLength } from "vuelidate/lib/validators";

export default {
  mixins: [validationMixin],

  validations: {
    sessionName: { required, maxLength: maxLength(20) }
  },

  props: {
    user: String
  },

  watch: {
    user: function() {
      this.userId = this.user;
    }
  },

  data: () => ({
    userId: "",
    sessionName: "",
    firstName: "",
    lastName: "",
    sessions: []
  }),

  computed: {
    sessionNameErrors() {
      const errors = [];
      if (!this.$v.sessionName.$dirty) return errors;
      !this.$v.sessionName.maxLength &&
        errors.push("Session name must be at most 20 characters long");
      !this.$v.sessionName.required && errors.push("Session name is required.");
      return errors;
    }
  },

  methods: {
    startSession() {
      console.log(`Start clicked: ${this.userId}`);
      this.$v.$touch();
      this.axios
        .post(
          `/user/${this.userId}/session_start`,
          JSON.stringify({
            name: this.sessionName
          })
        )
        .then(r => {
          console.log(r.data);
          const { firstName, lastName, timeTrackingSessions, id } = r.data;
          this.firstName = firstName;
          this.lastName = lastName;
          this.sessions = timeTrackingSessions;
          this.userId = id;

          this.$emit("sessionFormSubRes", this.$data);
        });
    }
  }
};
</script>
