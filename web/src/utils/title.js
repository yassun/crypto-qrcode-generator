export default {
  mounted() {
    let title = this.title
    if (title) {
      title = typeof title === 'function' ? title.call(this) : title
      document.title = `Crypto QR Code Generator - ${title}`
    }
  }
}
