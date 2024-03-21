<script>
/**

Questo codice è un componente Vue.js che visualizza il profilo di un utente. 
Ha una serie di proprietà come nickname, followStatus, followerCnt, followingCnt, postCnt, photos, following, followers ecc.,
e metodi come uploadFile, followClick, banClick, loadInfo.

Il componente utilizza la libreria Axios per interagire con un'API back-end per recuperare e aggiornare i dati. 
Ad esempio, il metodo followClick invia una richiesta PUT all'API per seguire l'utente o
una richiesta DELETE per smettere di seguire l'utente. 
Il metodo banClick invia una richiesta PUT per escludere l'utente o una richiesta DELETE per sbloccare l'utente.
Il metodo loadInfo recupera i dati del profilo dell'utente dall'API e aggiorna le proprietà del componente di conseguenza. 
Il metodo uploadFile consente all'utente di caricare una foto.

La proprietà calcolata currentPath restituisce l'id dell'utente dall'oggetto $route. 
La proprietà calcolata sameUser controlla se l'id dell'utente visualizzato è lo stesso dell'utente attualmente connesso. 
La proprietà watch controlla le modifiche nella proprietà currentPath e chiama il metodo loadInfo quando viene rilevata una modifica.

 */
export default {
  data: function () {
    return {
      errormsg: null,

      userExists: false,
      banStatus: false,

      nickname: "",
      user_id: 0,

      followStatus: false,
      currentIsBanned: false,

      followerCnt: 0,
      followingCnt: 0,
      postCnt: 0,

      photos: [],
      following: [],
    }
  },

  watch: {
    currentPath(newid, oldid) {
      if (newid !== oldid) {
        this.loadInfo()
      }
    },
  },

  computed: {
    currentPath() {
      return this.$route.params.id
    },

    sameUser() {
      return this.$route.params.id === localStorage.getItem('token')
    },
  },

  methods: {

    async uploadFile() {
      let fileInput = document.getElementById('fileUploader')
      const file = fileInput.files[0];
      const reader = new FileReader();
      reader.readAsArrayBuffer(file);
      reader.onload = async () => {
        // Post photo: /Users/:id/Photos
        let response = await this.$axios.post("/Users/" + this.$route.params.id + "/photos", reader.result, {
          headers: {
            'Content-Type': file.type
          },
        });
        //this.photos.append(response.data)
        this.postCnt += 1

        let response2 = await this.$axios.get("/Users/" + this.$route.params.id);
        this.photos = response2.data.posts != null ? response2.data.posts : []
      };
    },

    async followClick() {
      try {
        if (this.followStatus) {
          // Delete follow: /users/:id/followers/:follower_id
          await this.$axios.delete("/Users/" + localStorage.getItem('token') + "/followers/" + this.$route.params.id);
          this.followerCnt -= 1
        } else {
          // Put follow: /users/:id/followers/:follower_id
          await this.$axios.put("/Users/" + localStorage.getItem('token') + "/followers/" + this.$route.params.id);
          this.followerCnt += 1
        }
        this.followStatus = !this.followStatus
      } catch (e) {
        this.errormsg = e.toString();
      }

    },

    async banClick() {
      try {
        if (this.banStatus) {
          // Delete ban: /users/:id/banned_users/:banned_id
          await this.$axios.delete("/Users/" + localStorage.getItem('token') + "/banned_users/" + this.$route.params.id);
          this.loadInfo()
        } else {
          // Put ban: /users/:id/banned_users/:banned_id
          await this.$axios.put("/Users/" + localStorage.getItem('token') + "/banned_users/" + this.$route.params.id);
          this.followStatus = false
        }
        this.banStatus = !this.banStatus
      } catch (e) {
        this.errormsg = e.toString();
      }
    },

    async loadInfo() {
      if (this.$route.params.id === undefined) {
        return
      }
      try {
        // Get user profile: /Users/:id
        let response = await this.$axios.get("/Users/" + this.$route.params.id);
        this.banStatus = false

        this.userExists = true
        this.currentIsBanned = false
        if (response.status === 206) {
          this.banStatus = true
          return
        }
        if (response.status === 204) {
          this.userExists = false
        }
        this.nickname = response.data.nickname
        this.user_id = response.data.user_id
        this.followerCnt = response.data.n_follower
        this.followingCnt = response.data.n_following
        this.postCnt = response.data.posts != null ? response.data.posts.length : 0
        this.followStatus = response.data.followers != null ? response.data.followers.find(obj => obj.user_id === localStorage.getItem('token')) : false
        this.photos = response.data.posts != null ? response.data.posts : []
        this.followers = response.data.followers != null ? response.data.followers : []


      } catch (e) {
        this.currentIsBanned = true
      }
    },

    goToSettings() {
      this.$router.push(this.$route.params.id + '/settings')
    },

    removePhotoFromList(photo_id) {
      this.photos = this.photos.filter(item => item.photo_id !== photo_id)
      this.postCnt -= 1
    },
  },

  async mounted() {
    this.$axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('token')
    await this.loadInfo()
  },

}
</script>

<!--
  
È un singolo componente di file che visualizza le informazioni sul profilo di un utente, inclusi il nome utente, il numero di post,
il numero di follower e il numero di account che stanno seguendo.
Ha una sezione di intestazione che include il nome utente del profilo visualizzato e pulsanti per seguire/smettere di seguire 
e bloccare/sbloccare l'utente. Se l'utente corrente che sta visualizzando il profilo è lo stesso del profilo visualizzato, 
vedrà un'icona a forma di ingranaggio per accedere alle proprie impostazioni anziché i pulsanti Segui/Blocca.

La scheda profilo mostra le statistiche (conteggio post, conteggio follower, conteggio follower) del profilo solo se l'utente 
non è stato bannato.

Il profilo ha anche una sezione per visualizzare i post dell'utente, che è suddivisa in un'intestazione, un pulsante "Aggiungi" 
(se l'utente corrente è il proprietario del profilo) e l'elenco dei post. 
Se l'utente viene bannato viene visualizzato il messaggio "Non c'è ancora alcun post". 
Se l'utente ha post, viene visualizzata una griglia di componenti <Photo>. 
Ogni componente <Foto> visualizza una singola foto e le relative informazioni, come il numero di Mi piace e commenti.

-->

<template>

  <div class="profile-card d-flex flex-column align-items-center" v-if="!currentIsBanned && userExists">
    <div class="header">
      <h5 class="username">@{{ nickname }} </h5>
      <div class="actions">
        <template v-if="!sameUser">
          <button v-if="!banStatus" @click="followClick" class="btn btn-success">
            {{ followStatus ? "Non seguire più" : "Segui" }}
          </button>
          <button @click="banClick" class="btn btn-danger">
            {{ banStatus ? "Sblocca" : "Blocca" }}
          </button>
        </template>

        <template v-else>
          <button class="my-trnsp-btn" @click="goToSettings">
            <i class="my-nav-icon-gear fa-solid fa-gear"></i>
          </button>
        </template>
      </div>
    </div>
    <template v-if="!banStatus">
      <div class="stats mt-3 mb-3">
        <div class="post-count positioned-posts">Post: {{ postCnt }}</div>
        <div class="follower-count positioned-followers">Seguaci: {{ followerCnt }}</div>
        <div class="following-count positioned-following">Seguiti: {{ followingCnt }}</div>
      </div>
    </template>
  </div>

  <div class="profile-posts-container">
    <h2 class="posts-header text-center">Post</h2>

    <div style="padding-left: 150px; padding-right: 150px;">
      <div class="input-photo"  style="">
        <input id="fileUploader" type="file" accept=".png" >
        <button @click="uploadFile">Toccami</button>
      </div>
    </div>
    <hr class="posts-divider" />
    
    <div class="row" style="width: 100%;" v-if="!banStatus && postCnt > 0">
      <div v-for="(photo) in photos" class="col-sm-4" :key="photo.photo_id">
        <Photo :key="photo.photo_id" :owner="this.$route.params.id" :nickname="this.nickname" :photo_id="photo.photo_id"
          :comments="photo.comments" :likes="photo.likes" :upload_date="photo.timestamp" :isOwner="sameUser"
          @removePhoto="removePhotoFromList" />
      </div>
    </div>

    <div v-else class="mt-5" style="">
      <h2 class="no-posts-text">Non c'è ancora alcun post</h2>
    </div>

    <error-msg v-if="errormsg" :msg="errormsg"></error-msg>
  </div>
</template>

<style>

.input-photo{
  border: 1px solid black;
  border-radius: 5px;
  display: flex; 
  justify-content: center;
}

.posts-header {
  color: #333;
  font-weight: bold;
  margin-left: -50px;
  font-family: 'Times New Roman', Times, serif;
  font-size: 30px;

}

.posts-divider {
  border-top: 1px solid #333;
  margin: 20px 0;
}

.btn-add-photo {
  background-color: #333;
  color: #fff;
  border: none;
  padding: 10px 15px;
  border-radius: 50px;
  transition: all 0.3s ease-in-out;
  margin-left: -50px !important;
}

.btn-add-photo:hover {
  background-color: #fff;
  color: #333;
  cursor: pointer;
}

.no-posts-text {
  color: #333;
}

.positioned-posts {
  display: inline-block;
  margin-left: 40px;
  margin-right: 40px;
  text-align: center;
  font-family: 'Times New Roman', Times, serif;
  font-size: 25px;
}

.positioned-followers {
  display: inline-block;
  margin-left: 40px;
  margin-right: 40px;
  text-align: center;
  font-family: 'Times New Roman', Times, serif;
  font-size: 25px;
}

.positioned-following {
  display: inline-block;
  margin-left: 40px;
  margin-right: 40px;
  text-align: center;
  font-family: 'Times New Roman', Times, serif;
  font-size: 25px;
}

.my-nav-icon-gear {
  margin-left: 10px;
}

.header {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-left: -5px;
  margin-top: 95px;
  font-size: 15px;
  font-family: 'Times New Roman', Times, serif;
}

.actions {
  display: flex;
  justify-content: flex-end;
  margin-left: 30px;
}

.username {
  font-size: 30px;
  font-family: 'Times New Roman', Times, serif;
  color: black;

}

.col-sm-4 {
  margin-bottom: 0;
  margin-left: 55px;
  margin-right: -80px;
}

button.btn-success {
  font-size: 20px;
  font-family: 'Times New Roman', Times, serif;
}

button.btn-danger {
  font-size: 20px;
  font-family: 'Times New Roman', Times, serif;
}
</style>