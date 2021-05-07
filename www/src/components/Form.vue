<template>
  <form>
    <v-text-field
      v-model="firstName"
      :error-messages="firstNameErrors"
      :counter="20"
      label="First Name"
      required
      @input="$v.firstName.$touch()"
      @blur="$v.firstName.$touch()"
    ></v-text-field>
    <v-text-field
      v-model="lastName"
      :error-messages="lastNameErrors"
      :counter="20"
      label="Last Name"
      required
      @input="$v.lastName.$touch()"
      @blur="$v.lastName.$touch()"
    ></v-text-field>
    <v-btn class="mr-4" @click="submit">
      submit
    </v-btn>
    <v-btn @click="clear">
      clear
    </v-btn>
  </form>
</template>

<script>
import { validationMixin } from "vuelidate";
import { required, maxLength } from "vuelidate/lib/validators";

export default {
  mixins: [validationMixin],

  validations: {
    firstName: { required, maxLength: maxLength(20) },
    lastName: { required, maxLength: maxLength(20) }
  },

  data: () => ({
    id: "",
    firstName: "",
    lastName: "",
    timeTrackingSessions: []
  }),

  computed: {
    firstNameErrors() {
      const errors = [];
      if (!this.$v.firstName.$dirty) return errors;
      !this.$v.firstName.maxLength &&
        errors.push("First name must be at most 20 characters long");
      !this.$v.firstName.required && errors.push("First name is required.");
      return errors;
    },
    lastNameErrors() {
      const errors = [];
      if (!this.$v.lastName.$dirty) return errors;
      !this.$v.lastName.maxLength &&
        errors.push("Last name must be at most 20 characters long");
      !this.$v.lastName.required && errors.push("Last name is required.");
      return errors;
    }
  },

  methods: {
    submit() {
      this.$v.$touch();
      this.axios
        .post(
          "/user",
          JSON.stringify({
            firstName: this.firstName,
            lastName: this.lastName
          })
        )
        .then(r => {
          // console.log(r.data);
          const { firstName, lastName, timeTrackingSessions, id } = r.data;
          this.firstName = firstName;
          this.lastName = lastName;
          this.timeTrackingSessions = timeTrackingSessions;
          this.id = id;

          this.$emit("userFormSubRes", this.$data);
        });
    },
    clear() {
      this.$v.$reset();
      this.firstName = "";
      this.lastName = "";
    }
  }
};
</script>
