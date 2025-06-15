<template>
  <div>
    <section id="hero"
      class="hero w-100 h-100 p-3 mx-auto text-center d-flex justify-content-center align-items-center text-white">
      <main>
        <h4>Kepada <span>Bapak/Ibu/Saudara/i, </span></h4>
        <h1>Cat & Cut</h1>
        <p>Akan melangsungkan resepsi pernikahan dalam:</p>
        <div class="simply-countdown"></div>
        <a href="#home" class="btn btn-lg mt-4" @click="enableScroll">Lihat Undangan</a>
      </main>

    </section>

    <nav class="navbar navbar-expand-md bg-transparent sticky-top mynavbar">
      <div class="container">
        <a class="navbar-brand" href="#">Dino</a>
        <button class="navbar-toggler border-0" type="button" data-bs-toggle="offcanvas"
          data-bs-target="#offcanvasNavbar" aria-controls="offcanvasNavbar" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="offcanvas offcanvas-end" tabindex="-1" id="offcanvasNavbar" aria-labelledby="offcanvasNavbarLabel">
          <div class="offcanvas-header">
            <h5 class="offcanvas-title" id="offcanvasNavbarLabel">Dino</h5>
            <button type="button" class="btn-close" data-bs-dismiss="offcanvas" aria-label="Close"></button>
          </div>
          <div class="offcanvas-body">
            <div class="navbar-nav ms-auto">
              <a class="nav-link" href="#home">Home</a>
              <a class="nav-link" href="#info">Info</a>
              <a class="nav-link" href="#rsvp">RSVP</a>
              <a class="nav-link" href="#gifts">Gifts</a>
            </div>
          </div>
        </div>
      </div>
    </nav>

    <section id="home" class="home">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-md-8 text-center">
            <h2>Acara Pernikahan</h2>
            <h3>Diselenggarakan pada 20 November 2024 di Jakarta, Jawa Barat.</h3>
            <p>Oleh karena itu, d engan segala hormat, kami bermaksud untuk mengundang Bapak/Ibu, Saudara/i, untuk hadir
              pada acara pernikahan kami. </p>
          </div>
        </div>

        <div class="row couple">
          <div class="col-lg-6">
            <div class="row">
              <div class="col-8 text-end">
                <h3>Cat</h3>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Atque praesentium aut ipsa perferendis,
                  incidunt soluta?</p>
                <p>Putra dari Bpk. Lorem <br> dan <br> Ibu Ipsum</p>
              </div>
              <div class="col-4">
                <img :src="sandhika" class="/public/img-responsive rounded-circle">
              </div>
            </div>
          </div>

          <span class="heart"><i class="bi bi-heart-fill"></i></span>

          <div class="col-lg-6">
            <div class="row">
              <div class="col-4">
                <img :src="nova" class="/public/img-responsive rounded-circle">
              </div>
              <div class="col-8">
                <h3>Cut</h3>
                <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Atque praesentium aut ipsa perferendis,
                  incidunt soluta?</p>
                <p>Putra dari Bpk. Ipsum <br> dan <br> Ibu Lorem</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>


    <section id="info" class="info">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-md-8 col-10 text-center">
            <h2>Informasi Acara</h2>
            <p class="alamat">Alamat: Bambu Kuning <br></p>

            <iframe
              src="https://www.google.com/maps/embed?pb=!1m14!1m8!1m3!1d15843.359732941108!2d107.616864!3d-6.9097349!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x2e68e635c9284fc1%3A0x268bddf9bd0971c0!2sKologdam!5e0!3m2!1sen!2ssg!4v1693486483694!5m2!1sen!2ssg"
              width="100%" height="250" style="border:0;" allowfullscreen="" loading="lazy"
              referrerpolicy="no-referrer-when-downgrade"></iframe>

            <a href="https://goo.gl/maps/R9seweCadg9MNyox5" target="_blank" class="btn btn-light btn-sm my-3">Klik untuk
              membuka peta</a>
            <p class="description">Diharapkan untuk tidak salah alamat dan tanggal. Manakala tiba di tujuan namun tidak
              ada tanda-tanda sedang dilangsungkan pernikahan, boleh jadi Anda salah jadwal, atau salah tempat.</p>
          </div>
        </div>

        <div class="row justify-content-center mt-4">
          <div class="col-md-5 col-10">
            <div class="card text-center text-bg-light mb-5">
              <div class="card-header">Akad Nikah</div>
              <div class="card-body">
                <div class="row justify-content-center">
                  <div class="col-md-6">
                    <i class="bi bi-clock d-block"></i>
                    <span>08.00 - 10.00</span>
                  </div>
                  <div class="col-md-6">
                    <i class="bi bi-calendar3 d-block"></i>
                    <span>Minggu, 20 November 2023</span>
                  </div>
                </div>
              </div>
              <div class="card-footer">
                Saat acara akad diharapkan untuk kondusif menjaga kekhidmatan dan kekhusyuan seluruh prosesi.
              </div>
            </div>
          </div>
          <div class="col-md-5 col-10">
            <div class="card text-center text-bg-light">
              <div class="card-header">Resepsi</div>
              <div class="card-body">
                <div class="row justify-content-center">
                  <div class="col-md-6">
                    <i class="bi bi-clock d-block"></i>
                    <span>11.00 - selesai</span>
                  </div>
                  <div class="col-md-6">
                    <i class="bi bi-calendar3 d-block"></i>
                    <span>Minggu, 20 November 2023</span>
                  </div>
                </div>
              </div>
              <div class="card-footer">
                Saat acara akad diharapkan untuk kondusif menjaga kekhidmatan dan kekhusyuan seluruh prosesi.
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <section id="rsvp" class="rsvp">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-md-8 col-10 text-center">
            <h2>Konfirmasi Kehadiran</h2>
            <p v-if="!hasSubmitted">Isi form di bawah ini untuk melakukan konfirmasi kehadiran.</p>
            <p v-else class="text-success">Terima kasih telah mengkonfirmasi kehadiran Anda!</p>
          </div>
        </div>

        <!-- Display guest data for debugging -->
        <div class="text-center mb-3" v-if="guestData">
          <pre class="d-none">{{ guestData }}</pre>
        </div>

        <!-- Show different content based on attend status -->
        <div v-if="hasSubmitted">
          <div class="alert alert-info text-center">
            Anda sudah mengisi konfirmasi kehadiran:
            <strong>{{ guestData.attend ? 'Hadir' : 'Tidak Hadir' }}</strong>
          </div>

          <div class="text-center mt-3">
            <button @click="() => { hasSubmitted = false; showUpdateForm = true }" class="btn btn-outline-primary"
              v-if="!showUpdateForm">
              Update Konfirmasi
            </button>
          </div>
        </div>

        <!-- Form for new submission or update -->
        <form @submit.prevent="onSubmit" class="row row-cols-md-auto g-3 align-items-center justify-content-center"
          v-if="showUpdateForm">
          <div class="row col-12">
            <div class="col-4">
              <div class="mb-3">
                <label for="jumlah" class="form-label">Jumlah</label>
                <input type="number" class="form-control" id="jumlah" v-model="formData.jumlah" min="1" max="5"
                  required>
              </div>
            </div>
            <div class="col-8">
              <div class="mb-3">
                <label for="status" class="form-label">Konfirmasi</label>
                <select name="status" id="status" class="form-select" v-model="formData.status" required>
                  <option disabled value="">Pilih salah satu</option>
                  <option value="Hadir">Hadir</option>
                  <option value="Tidak Hadir">Tidak Hadir</option>
                </select>
              </div>
            </div>
            <div class="col-12">
              <div class="mb-3">
                <label for="message" class="form-label">Ucapan</label>
                <textarea class="form-control" id="message" v-model="formData.message" rows="3" required></textarea>
              </div>
            </div>
            <div class="col-12" style="margin-top: 35px;">
              <button type="submit" class="btn btn-primary col-12" :disabled="isSubmitting">
                {{ isSubmitting ? 'Mengirim...' : 'Kirim' }}
              </button>
              <button v-if="showUpdateForm" @click="showUpdateForm = false"
                class="btn btn-outline-secondary col-12 mt-2" type="button">
                Batal
              </button>
            </div>
          </div>
        </form>
      </div>
    </section>

    <section id="gifts" class="gifts">
      <div class="container">
        <div class="row justify-content-center">
          <div class="col-md-8 col-10 text-center">
            <span>ungkapan tanda kasih</span>
            <h2>Kirim Hadiah</h2>
            <p>Lorem ipsum dolor sit amet consectetur adipisicing elit. Magnam iure perferendis provident ab aliquam
              nemo.
            </p>
          </div>
        </div>

        <div class="row justify-content-center text-center">
          <div class="col-md-6">
            <ul class="list-group">
              <li class="list-group-item">
                <div class="fw-bold">BCA</div>
                123456789 - Cat
              </li>
              <li class="list-group-item">
                <div class="fw-bold">MANDIRI</div>
                987654321 - Cut
              </li>
              <li class="list-group-item">
                <div class="fw-bold">Saweria</div>
                <img :src="saweria" alt="Saweria QR" class="/public/img-thumbnail" width="150">
              </li>
            </ul>
          </div>
        </div>
      </div>
    </section>

    <footer>
      <div class="container">
        <div class="row">
          <div class="col text-center">
            <small class="block">&copy; 2025 All Rights Reserved.</small>
            <ul class="mt-3">
              <li><a href="#"><i class="bi bi-instagram"></i></a></li>
              <li><a href="#"><i class="bi bi-youtube"></i></a></li>
              <li><a href="#"><i class="bi bi-twitter"></i></a></li>
              <li><a href="#"><i class="bi bi-facebook"></i></a></li>
              <li><a href="#"><i class="bi bi-tiktok"></i></a></li>
            </ul>
          </div>
        </div>
      </div>
    </footer>

    <div id="audio-container">
      <audio id="song" loop ref="songRef">
        <source :src="musicUndangan" type="audio/mp3">
      </audio>

      <div class="audio-icon-wrapper" ref="audioIconWrapper" style="display: none" @click="playAudio">
        <i class="bi bi-disc"></i>
      </div>
    </div>
  </div>

</template>

<script setup>
import { onMounted, ref } from 'vue'
import sandhika from '@/assets/img/sandhika.png'
import nova from '@/assets/img/nofa.png'
import musicUndangan from '@/assets/save-and-sound.mp3'

const audioIconWrapper = ref(null)
const songRef = ref(null)
let isPlaying = false

function enableScroll() {
  window.onscroll = () => { }
  document.documentElement.style.scrollBehavior = 'smooth'
  playAudio()
}

function playAudio() {
  if (isPlaying) {
    song.pause()
    isPlaying = false
  } else {
    debugger
    const song = songRef.value
    const icon = audioIconWrapper.value.querySelector('i')
    if (song) {
      song.volume = 1
      audioIconWrapper.value.style.display = 'flex'
      song.play()
      isPlaying = true
      icon.classList.add('bi-disc')
      icon.classList.remove('bi-pause-circle')
    }
  }
}

async function fetchGuestData(guestId) {
  try {
    const response = await fetch(`http://localhost:8085/public/guest/${guestId}`);

    if (!response.ok) {
      throw new Error(`HTTP error! Status: ${response.status}`);
    }

    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Failed to fetch guest data:', error);
    throw error; // Re-throw to let caller handle
  }
}

async function updateLastView(guestId) {
  const url = `http://localhost:8085/public/guest/${guestId}`;

  try {
    const response = await fetch(url, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({}) // Empty object as per your curl example
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    // Handle empty response or JSON response
    const text = await response.text();
    return text ? JSON.parse(text) : { status: 'success' };
  } catch (error) {
    console.error('Error updating last view:', error);
    throw error;
  }
}




const guestData = ref({})

function getGuestIdFromUrl() {
  const urlParams = new URLSearchParams(window.location.search);
  const guestId = urlParams.get('guest_id');

  if (!guestId) {
    console.error('No guest_id parameter found in URL');
    return null;
  }

  return guestId;
}
const isSubmitting = ref(false);
const showUpdateForm = ref(false);
const hasSubmitted = ref(false);
const formData = ref({
  jumlah: 1,
  status: '',
  message: 'Happy Wedding :love'
});

// Your existing updateActivity function
async function updateActivity(guestId, message, attendStatus) {
  try {
    const response = await fetch(`http://localhost:8085/public/guest/message`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        id: guestId,
        message: message,
        attend: attendStatus === 'Hadir' // Convert to boolean
      })
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return await response.json();
  } catch (error) {
    console.error('Error updating activity:', error);
    throw error;
  }
}

// Form submission handler
const onSubmit = async () => {
  try {
    isSubmitting.value = true;

    // Validate form
    if (!formData.value.status) {
      alert('Silakan pilih status konfirmasi');
      return;
    }

    const guestID = getGuestIdFromUrl()
    await updateActivity(guestID, formData.value.message, formData.value.status)

    // Optional: Reset form after successful submission
    formData.value.status = '';
    hasSubmitted.value = true;
    showUpdateForm.value = false;

    debugger
  } catch (error) {
    console.error('Submission error:', error);
    alert(`Gagal mengirim data: ${error.message}`);
  } finally {
    isSubmitting.value = false;
  }
};

//
onMounted(async () => {
  let guestID = getGuestIdFromUrl()
  await updateLastView(guestID)
  let data = await fetchGuestData(guestID)
  guestData.value = data

  if (data.hasOwnProperty("attend")) {
    hasSubmitted.value = true;
    showUpdateForm.value = false;
  } else {
    hasSubmitted.value = false;
    showUpdateForm.value = true;
  }


  window.scrollTo(0, 0)
  if (window.location.hash !== '') {
    history.replaceState(null, '', window.location.pathname + window.location.search)
  }
  if (!document.querySelector('link[href="https://fonts.gstatic.com"]')) {
    const preconnectFontsGstatic = document.createElement('link')
    preconnectFontsGstatic.rel = 'preconnect'
    preconnectFontsGstatic.href = 'https://fonts.gstatic.com'
    preconnectFontsGstatic.crossOrigin = ''
    document.head.appendChild(preconnectFontsGstatic)
  }

  // Google Fonts stylesheet
  if (!document.querySelector('link[href^="https://fonts.googleapis.com/css2?family=Sacramento"]')) {
    const googleFonts = document.createElement('link')
    googleFonts.rel = 'stylesheet'
    googleFonts.href =
      'https://fonts.googleapis.com/css2?family=Sacramento&family=Work+Sans:wght@100;300;400;600;700&display=swap'
    document.head.appendChild(googleFonts)
  }

  // Bootstrap Icons CSS
  if (!document.querySelector('link[href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css"]')) {
    const bootstrapIcons = document.createElement('link')
    bootstrapIcons.rel = 'stylesheet'
    bootstrapIcons.href = 'https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css'
    document.head.appendChild(bootstrapIcons)
  }

  // Bootstrap Icons CSS
  if (!document.querySelector('link[href="/style.css')) {
    const bootstrapIcons = document.createElement('link')
    bootstrapIcons.rel = 'stylesheet'
    bootstrapIcons.href = 'style.css'
    document.head.appendChild(bootstrapIcons)
  }

  // Bootstrap Icons CSS
  if (!document.querySelector('link[href="/countdown/simplyCountdown.theme.default.css')) {
    const bootstrapIcons = document.createElement('link')
    bootstrapIcons.rel = 'stylesheet'
    bootstrapIcons.href = 'countdown/simplyCountdown.theme.default.css'
    document.head.appendChild(bootstrapIcons)
  }

  // simplyCountdown external JS only
  if (!document.querySelector('script[src="countdown/simplyCountdown.min.js"]')) {
    const simplyCountdownScript = document.createElement('script')
    simplyCountdownScript.src = 'countdown/simplyCountdown.min.js' // local file, but included in your original, so omit or include?
    simplyCountdownScript.defer = true
    document.head.appendChild(simplyCountdownScript)
    simplyCountdownScript.onload = () => {
      console.log("SimplyCountdown loaded")

      if (window.simplyCountdown) {
        window.simplyCountdown('.simply-countdown', {
          year: 2025,
          month: 7,
          day: 10,
          hours: 8,
          words: {
            days: { singular: 'hari', plural: 'hari' },
            hours: { singular: 'jam', plural: 'jam' },
            minutes: { singular: 'menit', plural: 'menit' },
            seconds: { singular: 'detik', plural: 'detik' }
          },
        })
      } else {
        console.error("window.simplyCountdown not defined")
      }
    }
  }
})
</script>

<style scoped>
@import "https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css";
@import "https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.5/font/bootstrap-icons.css";
</style>
