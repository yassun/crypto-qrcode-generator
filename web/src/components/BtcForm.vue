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
            :rules="amountRules"
            label="Amount (optional)"
            type="number"
          ></v-text-field>

          <v-text-field
            v-model="qrlabel"
            :rules="qrlabelRules"
            label="Label (optional): Luke-Jr..."
          ></v-text-field>

          <v-textarea
            v-model="message"
            :rules="messageRules"
            label="Message (optional): Donation for project..."
            value=""
          ></v-textarea>

         <v-row align="center" justify="center">
          <v-btn
            color="success"
            @click="submit"
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
    url:"http://localhost:8000/generate-qr/btc",
    address: "",
    addressRules: [
        v => !!v || 'Address is required',
        v => (v && v.length <= 100) || 'Address must be less than 100 characters',
    ],
    amount: "",
    amountRules: [
        v => (v.length <= 100) || 'Amount must be less than 100 characters',
    ],
    qrlabel: "",
    qrlabelRules: [
        v => (v.length <= 255) || 'Label must be less than 255 characters',
    ],
    message: "",
    messageRules: [
        v => (v.length <= 255) || 'Message must be less than 255 characters',
    ],

  }),
  methods: {
     submit () {
       if (!this.$refs.form.validate()) {
         return
       }
       let params = new FormData()
       params.append('address', this.address)
       params.append('amount', this.amount)
       params.append('label', this.label)
       params.append('message', this.message)
       this.$axios.post(this.url, params)
         .then(function(res){
             console.log(res)
         })
         .catch(function(res){
             console.log(res)
         })
     },
  }
};
</script>
