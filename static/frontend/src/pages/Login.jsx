import React, { Component } from "react";
import { Link } from "react-router-dom";
import { Redirect } from "react-router-dom";
import { isset } from "../js/api"

class Login extends Component {
    constructor(props) {
        super(props);

        this.state = {
            username: '',
            password: '',
            redirectToReferrer: false
        }
    }

    onChange = (e) => {
        this.setState({
            [e.target.name]: e.target.value
        });
    }

    

        gotoURL(){
            let url = `challenges`;
            document.location=(url);
        }

    handleSubmit = async (event) => {
        event.preventDefault();
            
        var data = JSON.stringify({ 
            "user": this.state.username, 
            "pass": this.state.password
        });

        // var url="http://localhost:5431/api/login"
        var url="/api/login"        
        const response = await fetch(url, {
            method:"POST",
            body: data,
            mode: "cors",
            credentials: "include",
        });

        const json = await response.json();
        if (json.Error === "ok") {
            await isset();
            localStorage.setItem("user", "konnichiwa UwU");
            this.setState(() => ({
                redirectToReferrer: true
            }))

        } else {
            alert(json.Error)
        }

    };

    render() {
        
        if (this.state.redirectToReferrer === true) {
            this.gotoURL();
        }


        return (
            <React.Fragment>
                <div className="container">
                    <div className="row align-items-center h-100">
                        <div className="col-md-5 mx-auto">
                            <div className="bg-transparent mt-5">
                                <h2>Login</h2>
                                <hr />
                                <form name="signin">
                                    <div className="form-group">
                                        <input name="username" value={this.state.username} type="text" className="form-control" placeholder="Your Email *"
                                            required onChange={this.onChange} />
                                    </div>
                                    <div className="form-group">
                                        <input name="password" type="password" className="form-control"
                                            placeholder="Your Password *" required onChange={this.onChange} />
                                    </div>
                                    <div className="form-group">
                                        <input type="button" className="btn btn-success btn-block" value="Login" onClick={this.handleSubmit} required />
                                    </div>
                                    <div className="form-group text-center">
                                        {/* <a href="#">Forget Password? </a> */}
                                        <Link to={`/register`}>Not registered?</Link>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </React.Fragment>
        )
    }
}

export default Login;
