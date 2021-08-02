import React, { Component } from "react";
import { validateContact } from '../js/helper.js';
import { api } from  "../js/api.js"

class settings extends Component {
    constructor(props) {
        super(props);
        this.state = {
            old: {
                fullname: '',
                email: '',
                age: '',
                institution: '',
                contact: '',
                place: '',
                district: '',
                state: ''
            },
            oldpwd: '',
            newpwda: '',
            newpwdb: '',
			isLoading: true
        }
    }

    async componentDidMount() {
        const query = `
        query fullyuser($in: String!) {
            userdata(id: $in) {
              username
              age
              fullname
              email
              gender
            }
          }
        `
        const newLocal = await api("fullyuser", { in: "me" }, query);
        // alert(newLocal.userdata)
        var response = newLocal.userdata;
        // .then((response) => {
        //     console.log(response);
        this.setState({
            old: {
                fullname: response.fullname,
                username: response.username,
                email: response.email,
                age: response.age,
                gender: response.gender,
            },
            fullname: response.fullname,
            username: response.username,
            email: response.email,
            age: response.age,
            gender: response.gender,
        });
        this.setState({
			isLoading: false
		});
    }
    onChange = (e) => {
        this.setState({
            [e.target.name]: e.target.value
        });
    }

    handlePasswordSubmit = async (event) => {
        event.preventDefault();
        let creds = {
            oldpwd: this.state.oldpwd,
            newpwda: this.state.newpwda,
            newpwdb: this.state.newpwdb,
        }

        let query = `mutation resers($in: resetpwd){reset_pwd(input: $in)}`
        let req = await api("resers", { "in": creds }, query);
        let resp = req.reset_pwd;
        alert(resp)
    }

    handleProfileSubmit = async (event) => {
        event.preventDefault();
        
        if (this.state.fullname==="" || this.state.Age===""   || this.state.gender===""){
            alert("Empty fields spotted");
            return false;
        }
        let data = {
            fullname: this.state.fullname,
            age: this.state.age,
            gender: this.state.gender,
        }

        /* 
            Check if user has made any changes in fields. 
            Send request only if changed.
        */
        let sendRequest = false;

        for (var key in this.state.old) {
            if (this.state.old.hasOwnProperty(key)) {
                if (this.state[key] != this.state.old[key]) {
                    sendRequest = true;
                    break;
                }
            }
        }

        if (sendRequest === true) {
            const query = `mutation usrupdate($in: useredit){ updateUser(input: $in)}`
            var req = await api("usrupdate", { "in": data }, query)
            if (req.updateUser==='OK'){
                this.state.old= {
                        fullname: this.state.fullname,
                        age: this.state.age,
                        institution: this.state.institution,
                        gender: this.state.gender,
                }
                alert('ok')
            } else {
                alert(req.updateUser)
            }
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
        return (
            <div className="container h-100">
                <div className="row h-90">
                    <div className="col-md-6 mx-auto mt-5">
                        <h2>Settings</h2>
                        <hr />
                        <form name="register-data" onSubmit={this.handleProfileSubmit}>
                            <div className="form-group">
                                <input name="fullname" type="text" onChange={this.onChange} className="form-control" placeholder="Your Name *" value={this.state.fullname} required />
                            </div>
                            <div className="form-group">
                                <input name="username" type="text" className="form-control" placeholder="Your Username *" value={this.state.username} required disabled />
                            </div>
                            <div className="form-group">
                                <input name="email" type="text" className="form-control" placeholder="Your Email *" value={this.state.email} required disabled />
                            </div>
                            <div className="form-group">
                                <input name="age" type="text" onChange={this.onChange} className="form-control" placeholder="Age *" value={this.state.age} required />
                            </div>
                            <div className="form-group">
                                <input name="gender" type="text" onChange={this.onChange} className="form-control" placeholder="Gender *" value={this.state.gender} required />
                            </div>
                            <div className="form-group">
                                <input type="submit" className="btn btn-success btn-block" value="Update" required />
                            </div>
                        </form>
                    </div>
                    <div className="col-md-5 mx-auto mt-5">
                        <h2>Security</h2>
                        <hr />
                        <form onSubmit={this.handlePasswordSubmit}>
                            <div className="form-group">
                                <input name="oldpwd" type="password" className="form-control" placeholder="Current Password *" required onChange={this.onChange} />
                            </div>
                            <div className="form-group">
                                <input name="newpwda" type="password" className="form-control" placeholder="New Password *" required onChange={this.onChange} />
                            </div>
                            <div className="form-group">
                                <input name="newpwdb" type="password" className="form-control" placeholder="Confirm New Password *" required onChange={this.onChange} />
                            </div>
                            <div className="form-group">
                                <input type="submit" className="btn btn-success btn-block" value="Update" required />
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        )
    }
}

export default settings;