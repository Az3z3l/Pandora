import React, { Component } from "react";
import { validateContact, validateEmail } from '../js/helper.js';
import { Link } from "react-router-dom";


class Register extends Component {

    constructor(props) {
        super(props);

        this.state = {
            fullname: '',
            username: '',
            email: '',
            password1: '',
            password2: '',
            age: '',
            gender: '',
        }

        // this.handleInputChange = this.handleInputChange.bind(this);
    }

    // onChange = (e) => {
    //     this.setState({
    //         [e.target.name]: e.target.value
    //     });
    // }

    handleInputChange = (event) => {
        this.setState({ 
            [event.target.name]: event.target.value 
        })
    }
    

    // onFileChange = async (event) => { 
    //     // Update the state 
    //     let idcar = event.target.files[0]

    //     if (idcar) {
    //         const reader = new FileReader();
    //         reader.onload = this._handleReaderLoaded.bind(this)
    //         reader.readAsDataURL(idcar)
    //     }
    // }; 

    // _handleReaderLoaded = (readerEvt) => {
    //     let binstring = readerEvt.target.result
    //     let image = btoa(String(binstring))
    //     this.setState({
    //         file: image
    //     })
    // }

    gotoURL(){
        let url = `/login`;
        this.props.history.push(url);
    }


    handleSubmit = async (event) => {
        event.preventDefault();
        if (this.state.email === "" || this.state.fullname === "" || this.state.password1 === "" || this.state.password2 === "") {
            alert("All fields must be filled out");
            return false;
        }

        if (this.state.password1 !== this.state.password2){
            alert("password not matching")
            return
        }

        if (this.state.password1.length < 8){
            alert("Length of Password must be greater than 8")
            return
        }


        if (!validateEmail(this.state.email)) {
            alert("Not a proper email");
            return false;
        }
        
        var data = JSON.stringify({ 
            "email": this.state.email,
            "uname": this.state.username,
            "fullname": this.state.fullname,
            "pword1": this.state.password1,
            "pword2": this.state.password2,

        });

        // var url="http://localhost:5431/api/register"        
        
        var url="/api/register"
        const response = await fetch(url, {
            method:"POST",
            body: data,
        });

        const json = await response.json();

        if (json.Error==='Successful'){
            alert("Registration successful. Login to continue")
            return this.gotoURL();
        }
        else {
            alert(json.Error)
        }

    }


    render() {
        return (
            <React.Fragment>
                <div className="container h-100">
                    <div className="row align-items-center h-90">
                        <div className="col-md-5 mx-auto mt-5">
                            <h2>Register</h2>
                            <hr />
                            <form name="register-data" onSubmit={this.handleSubmit}>
                                <div className="form-group">
                                    <input name="fullname" value={this.state.fullname} type="text" className="form-control" placeholder="Your Name *"  required onChange={this.handleInputChange}/>
                                </div>
                                <div className="form-group">
                                    <input name="username" value={this.state.username} type="text" className="form-control" placeholder="Your Username "  onChange={this.handleInputChange}/>
                                </div>
                                <div className="form-group">
                                    <input name="email" value={this.state.email} type="text" className="form-control" placeholder="Your Email *"  required onChange={this.handleInputChange}/>
                                </div>
                                  <div className="form-group">
                                    <input name="password1"  type="password" className="form-control" placeholder="Your Password *"  required onChange={this.handleInputChange}/>
                                </div>
                                <div className="form-group">
                                    <input name="password2" type="password" className="form-control" placeholder="Your Password again *"  required onChange={this.handleInputChange}/>
                                </div>
                                <div className="form-group">
                                    <input type="submit" className="btn btn-success btn-block" value="Register" required/>
                                </div>
                                <div className="form-group">
                                    <Link className="btn btn-block text-white" to={`/login`}>Already Registered?</Link>
                                </div>
                            </form>
                        </div>
                    </div>
                </div>
            </React.Fragment>
        )
    }
}

export default Register;