export default {
    template: `  
    <div class="row">
      <div class="position-absolute top-0 end-0 p-3" style="z-index: 5" :class="{'d-none': !messageToast, 'd-block': messageToast}">
        <div class="toast align-items-center text-white border-0 show" 
        :class="{'bg-primary' : message.status, 'bg-danger' : !message.status}">

        <div class="d-flex">
          <div class="toast-body">
            {{message.content}}
          </div>
          <button type="button" class="btn-close btn-close-white me-2 m-auto" @click="messageToast = false"></button>
        </div>
        </div>
      </div>

      <h1 class="mx-auto">RegisterSevice</h1>
      <div class="col-lg-3">
        <label for="ID" class="form-label">ID</label>
        <input class="form-control" v-model="id">
      </div>
      <div class="col-lg-3">
        <label for="port" class="form-label">Port</label>
        <input class="form-control" v-model="port">
      </div>
      <div class="col-lg-3">
        <label for="port" class="form-label">Service</label>
        <input class="form-control" v-model="service">
      </div>
      <div class="col-lg-3 mt-auto">
        <button class="btn btn-primary" @click="registerService">Register</button>
      </div>
      <table class="table table-striped">
        <thead>
          <tr>
            <th scope="col">Uid</th>
            <th scope="col">Port</th>
            <th scope="col">Service</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="{index, uid, port, service} of hosted" :key="index">
            <td>{{uid}}</td>
            <td>{{port}}</td>
            <td>{{service}}</td>
          </tr>
        </tbody>
      </table>
    </div>`,
    data: function(){
      return {
        id: '',
        port: '',
        service: '',
        message: {},
        messageToast: false,
        hosted: []
      }
    },
    methods: {
      registerService: async function(){
        let result = await fetch(`http://140.134.26.99:8001/RegisterService/${this.id}/${this.port}/${this.service}`,{
            method: 'GET',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
        }).then(resp => resp.json())
        this.message.content = result.message;
        this.message.status = result.status === 'Error' ? false : true;
        this.messageToast = true;
        this.loadHost();
        setTimeout(() => {
          this.messageToast = false;
        }, 10000)
      },
      loadHost: async function(){
        this.hosted.splice(0, this.hosted.length);
        let result = await fetch(`http://140.134.26.99:8001/Loadservice`,{
          method: 'GET',
          headers: {
              'Content-Type': 'application/x-www-form-urlencoded'
          },
        }).then(resp => resp.json())
        let host = JSON.parse(atob(result.json));
        let index = 1;
        host.forEach(element => {
          this.hosted.push({
            index: index,
            uid: element.Uid,
            port: element.Port,
            service: element.Servicetype
          })
          index ++;
        });
      }
    },
    mounted: async function(){
      this.loadHost()
    }
  }
  
  