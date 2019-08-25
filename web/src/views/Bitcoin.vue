<template>
<v-container>
  <v-layout justify-center>
    <v-flex xs12 md6>
      <v-card class="pa-4">

        <v-row align="center" justify="center">
          <v-img
            :src="qrcodeImage"
            aspect-ratio="1"
            class="grey lighten-2"
            max-width="175"
            max-height="175"
          ></v-img>
        </v-row>

        <v-row align="center" justify="center" class="pt-2">
          <v-btn
            @click="downloadQRCode"
            v-bind:disabled="isDownload"
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
            @click="generateQRCode"
          >
            Generate QR Code
          </v-btn>
          </v-row>
        </v-form>
      </v-card>

    </v-flex>
  </v-layout>
  <v-snackbar
    color="green darken-1"
    v-model="okSnackbar"
    :timeout=2000
  >
  <v-icon>mdi-check-circle</v-icon>QR code generation is complete.
    <v-btn
      color="blue"
      text
      @click="okSnackbar = false"
    >
      Close
    </v-btn>
  </v-snackbar>
  <v-snackbar
    color="red accent-1"
    v-model="ngSnackbar"
    :timeout=2000
  >
  <v-icon>mdi-alert-circle</v-icon>QR code generation failed.
    <v-btn
      color="blue"
      text
      @click="ngSnackbar = false"
    >
      Close
    </v-btn>
  </v-snackbar>

</v-container>
</template>

<script>
import SampleQrImage from "@/assets/sampleqr.png";
export default {
  data: () => ({
    title: 'Bitcoin',
    isDownload: true,
    okSnackbar: false,
    ngSnackbar: false,
    qrcodeImage: SampleQrImage,
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
    downloadQRCode () {
      let url = this.qrcodeImage.replace(/^data:image\/[^;]+/,
          'data:application/octet-stream')
      window.open(url)
    },
    generateQRCode () {
      if (!this.isValidForm()) {
        return
      }
     this.reqBtcGenerate()
    },
    btcReqParams(){
      let params = new FormData()
      params.append('address', this.address)
      params.append('amount', this.amount)
      params.append('label', this.label)
      params.append('message', this.message)
      return params
    },
    isValidForm(){
      return this.$refs.form.validate()
    },

    reqBtcGenerate(){
      let url = "http://localhost:8000/generate-qr/btc"
      let self = this
      this.$axios.post(url, this.btcReqParams())
        .then(function(res){
            self.okSnackbar = true
            self.isDownload = false
            self.qrcodeImage = 'data:image/jpeg;base64,' +  res["data"]["qr"]
        })
        .catch(function(res){
            self.ngSnackbar = true
            console.log(res)
        })
    }

  }
};
</script>
