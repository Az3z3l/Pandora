import React, { Component } from "react";
import { api } from  "../../js/api.js"
import Button from 'react-bootstrap/Button'


class Admineditnotification extends Component {

    constructor(props) {
        super(props);

        const { id } = this.props.match.params

        this.state = {
            ID: id,
            title: '',
            description: '',
            unavailable: false,
            isLoading: true,
        }

    }
    
    async componentDidMount() {


        const query = `
        query oneNote($in:String!){
            onenotify(id: $in){
              name
              description
            }
          }
        `
        let data = {
            in: this.state.ID
        }
        const newLocal = await api("oneNote", data, query);
        var response = newLocal.onenotify;
        if (response === null){
            this.setState({
                unavailable: true,
            })
        } else {
            this.setState({

                title: response.name,
                description: response.description,

            });

        }
        this.setState({
			isLoading: false
		});
    }

    handleInputChange = (event) => {
        this.setState({ 
            [event.target.name]: event.target.value 
        })
    }

    notificationHandle = async (event) =>{
        event.preventDefault();
        var data = {
            id: this.state.ID,
            name: this.state.title,
            description : this.state.description
        }
        let query = `mutation editnot($in:notifiedit!){
            edit_notification(input:$in)  
        }`

        let req = await api("editnot", { "in": data }, query);
        alert(req.edit_notification)
    }

    handleDelete = async (event) => {
        let z = (window.confirm('You sure you wanna delete this notification?'))
        if (!z){
            return
        }
        let operationName = "delnot";
		let variables = {
			"in": this.state.ID,
		};
		let query = `
		mutation delnot($in: String!){
            delete_notification(id: $in)
          }
        `
        var a = await api(operationName, variables, query)
        var msg = a.delete_notification;
        if (msg === "OK"){
            document.location="../";
        }
        else{
            alert(msg);
            document.location="../";
        }
    }


    render() {
        if (this.state.isLoading) {
			return (
				<React.Fragment>
					<div class="spinner"></div>
				</React.Fragment>
			);
        }
        
        if (this.state.unavailable) {
            return (
				<React.Fragment>
					<div className="container">
						<div className="row mt-5">
							<h2>Edit Notification</h2>
						</div>
						<br />
						<h3>Notification not available</h3>
					</div>
				</React.Fragment>
			);
        }
        else{
            return (
                <div class="container h-100">
                    <div class="row align-items-center h-90">
                        <div class="col-md-5 mx-auto mt-5">
                            <h2>Edit Notification</h2>
                            <hr />
                            <form name="register-data" onSubmit={this.notificationHandle}>
                                <div class="form-group">
                                    <input name="title" type="text" class="form-control" placeholder="Notification title *" value={this.state.title}  onChange={this.handleInputChange} required />
                                </div>
                                <div class="form-group">
                                    <textarea name="description" type="text" class="form-control" placeholder="Notification description *" value={this.state.description} onChange={this.handleInputChange} required />
                                </div>
                                
                                <div class="form-group">
                                    <input type="submit" class="btn btn-success btn-block" value="Edit Notification" required/>
                                </div>
                                <div class="form-group">
                                    <Button className="btn btn-block" variant="danger" onClick={this.handleDelete}>Delete Notification</Button>
                                </div>
                                
                            </form>
                        </div>
                    </div>
                </div>
            )
        }
    }
}

export default Admineditnotification;
