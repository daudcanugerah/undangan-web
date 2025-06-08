import { defineStore } from 'pinia'
import { reactive } from 'vue'


// https://example.com/undangan/:userTemplate_id/invit/:guest_id
const userTemplateManagerList = [
  {
    id: 1,
    user_id: 2,
    base_userTemplate_id: 1,
    state: 1,
    slug: "AL2U91KK",
    url: "http://localhost:8085/undangan-ku/xyzdaud",
    message_template: {
      whatsapp: {
        text: `
✨ UNDANGAN RESEPSI PERNIKAHAN ✨

Kepada Yth.
Bapak/Ibu/Saudara/i *{{name}}*
di tempat
{{address}}

Tanpa mengurangi rasa hormat, perkenankan kami mengundang Bapak/Ibu/Saudara/i untuk menghadiri acara Resepsi Pernikahan kami:

💑 *You & Me*

📅 *Hari/Tanggal:* Sabtu, 23 November 2024  
⏰ *Pukul:* 19.00 WIB - 21.00 WIB  
📍 *Tempat:*  
Masjid Salman Al-Farisi  
Komp Bulog, Jl. H. Ten Raya No. 14 RT 14/ RW 7  
Kayu Putih, Kec. Pulo Gadung  
Kota Jakarta Timur, DKI Jakarta 13210

🔗 Info lengkap & RSVP:  
{{url}}

Merupakan suatu kebahagiaan bagi kami apabila Bapak/Ibu/Saudara/i berkenan hadir dan memberikan doa restu.

🙏 *Kami yang berbahagia*  
You & Me

📩 Mohon maaf undangan hanya dikirim via pesan ini karena keterbatasan jarak dan waktu.
`
      },
      email: {
        text: `
Kepada Yth.  
Bapak/Ibu/Saudara/i {{name}}  
di tempat

Dengan hormat,

Tanpa mengurangi rasa hormat, izinkan kami mengundang Bapak/Ibu/Saudara/i untuk menghadiri acara Resepsi Pernikahan kami:

You & Me

Hari/Tanggal: Sabtu, 23 November 2024  
Waktu: 19.00 WIB – 21.00 WIB  
Tempat:  
Masjid Salman Al-Farisi  
Komp Bulog, Jl. H. Ten Raya No. 14 RT 14/ RW 7  
Kayu Putih, Kec. Pulo Gadung  
Kota Jakarta Timur, DKI Jakarta 13210

Informasi lengkap dapat diakses melalui tautan berikut:  
👉 https://temanhidupku.webinvit.id/zahra-daud/QN2O60

Kehadiran dan doa restu dari Bapak/Ibu/Saudara/i merupakan kehormatan dan kebahagiaan bagi kami.

Hormat kami,  
You & Me

*Mohon maaf undangan hanya dibagikan melalui pesan ini karena keterbatasan jarak dan waktu.*
`,
      },
      telegram: {
        text: `
✨ *UNDANGAN RESEPSI PERNIKAHAN* ✨

Kepada Yth.  
Bapak/Ibu/Saudara/i *{{name}}*  
_di tempat_

Tanpa mengurangi rasa hormat, perkenankan kami mengundang Bapak/Ibu/Saudara/i untuk menghadiri acara Resepsi Pernikahan kami:

💑 *You & Me*

📅 *Hari/Tanggal:* Sabtu, 23 November 2024  
⏰ *Pukul:* 19.00 WIB – 21.00 WIB  
📍 *Tempat:*  
Masjid Salman Al-Farisi  
Komp Bulog, Jl. H. Ten Raya No. 14 RT 14/ RW 7  
Kayu Putih, Pulo Gadung, Jakarta Timur

🔗 *Info & RSVP:*  
[Klik di sini](https://temanhidupku.webinvit.id/zahra-daud/QN2O60)

Merupakan suatu kebahagiaan bagi kami apabila Bapak/Ibu/Saudara/i berkenan hadir dan memberikan doa restu.

🙏 *Kami yang berbahagia*  
You & Me

📩 _Mohon maaf undangan hanya dibagikan melalui pesan ini karena keterbatasan jarak dan waktu._
`
      },
      sms: {
        text: `Yth. {{name}}, kami mengundang ke Resepsi You & Me, Sabtu 23 Nov 2024, 19.00 WIB, Masjid Salman Al-Farisi, Jakarta. Info: temanhidupku.webinvit.id/zahra-daud/QN2O60`,
      },
    },
    name: 'Golden Wedding Classic',
    cover_image: "https://wevitation.com/img/slider/slide-1.jpg",
    created_at: new Date('2025-05-30T19:25:00Z'),
    updated_at: new Date('2025-05-30T19:25:00Z'),
    expire_at: new Date('2025-07-30T19:25:00Z'),
  },
]

export const useUserTemplateManagerStore = defineStore('userTemplateManager', {
  state: () => ({
    list: [],
  }),
  getters: {
    getList(state) {
      return state.list
    },
    getTemplate(state) {
      return (id) => {
        return state.list.filter(e => e.id == id)[0] || {}
      }
    }
  },
  actions: {
    fetch() {
      this.list = userTemplateManagerList.map(e => ({
        ...e,
      }))
    },
    delete(id = "") {
      this.list = this.list.filter(user => user.id !== id);
    },
    create(data = {}) {
      this.list.push(data);
    },
    update(id, data = {}) {
      const index = this.list.findIndex(user => user.id === id);
      if (index !== -1) {
        this.list[index] = { ...this.list[index], ...data };
      }
    },
  },
});
