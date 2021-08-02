import React, { Component } from "react";
import { api } from  "../../js/api.js"


class Adminaddnotification extends Component {

    constructor(props) {
        super(props);

        const { challid } = this.props.match.params

        this.state = {
            title: '',
            description: '',
        }

    }      
    handleInputChange = (event) => {
        this.setState({ 
            [event.target.name]: event.target.value 
        })
    }

    notificationHandle = async (event) =>{
        event.preventDefault();
        var data = {
            name: this.state.title,
            description : this.state.description
        }
        let query = `mutation Addnotify($in:notificationinp!){
            add_notifications(input: $in)
          }`

        let req = await api("Addnotify", { "in": data }, query);
        alert(req.add_notifications)
    }

    render() {
        return (
            <div class="container h-100">
                <div class="row align-items-center h-90">
                    <div class="col-md-5 mx-auto mt-5">
                        <h2>Add Notification</h2>
                        <hr />
                        <form name="register-data" onSubmit={this.notificationHandle}>
                            <div class="form-group">
                                <input name="title" type="text" class="form-control" placeholder="Notification title *" value={this.state.title}  onChange={this.handleInputChange} required />
                            </div>
                            <div class="form-group">
                                <input name="description" type="text" class="form-control" placeholder="Notification description *" value={this.state.description} onChange={this.handleInputChange} required />
                            </div>
                            
                            <div class="form-group">
                                <input type="submit" class="btn btn-success btn-block" value="Add Notification" required/>
                            </div>
                            
                        </form>
                    </div>
                </div>
            </div>
        )
    }
}

export default Adminaddnotification;
