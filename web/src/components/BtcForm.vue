<template>
<v-container>
  <v-layout justify-center>
    <v-flex xs12 md6>
      <v-card class="pa-4">

        <v-row align="center" justify="center">
          <v-img
            src="https://placehold.jp/175x175.png"
            aspect-ratio="1"
            class="grey lighten-2"
            max-width="175"
            max-height="175"
          ></v-img>
        </v-row>

        <v-row align="center" justify="center" class="pt-2">
          <v-btn
            fab
            small
            right
            top
          >
            <v-icon>mdi-download</v-icon>
          </v-btn>
        </v-row>

        <v-form
            ref="form"
            v-model="valid"
            lazy-validation
          >

          <v-text-field
            v-model="address"
            :rules="addressRules"
            label="BTC Address: 17VZNX1SN5Nt........"
            required
          ></v-text-field>

          <v-text-field
            v-model="amount"
            :rules="amountlRules"
            label="Send amount"
            required
          ></v-text-field>

          <v-text-field
            v-model="label"
            :rules="labelRules"
            label="Label (optional): Luke-Jr..."
          ></v-text-field>

          <v-textarea
            name="input-7-1"
            label="Message (optional): Donation for project..."
            value=""
            height="60"
          ></v-textarea>

         <v-row align="center" justify="center">
          <v-btn
            color="success"
            @click="validate"
          >
            Generate QR Code
          </v-btn>
          </v-row>
        </v-form>
      </v-card>
    </v-flex>
  </v-layout>
</v-container>
</template>

<script>
export default {
  data: () => ({
  }),
};
</script>
