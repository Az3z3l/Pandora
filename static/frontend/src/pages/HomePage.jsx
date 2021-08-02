import React, { Component } from "react";
import HomePageLogo from '../images/d2.png';
import { Link } from "react-router-dom";
import Particles from 'react-particles-js';


class HomePage extends Component {
    render() {
        let buttons;
        if(localStorage.getItem('user') && (document.cookie.indexOf('auth')!==-1)){
            
            } else {
                buttons = 
                <p className="lead">
                    <Link className="btn btn-primary btn-lg" to={`/login`}>
                        Login
                    </Link>
                    <Link className="btn btn-primary btn-lg" to={`/register`}>
                        Register
                    </Link>    
                </p>
            }
        return (
            <React.Fragment>
                
            {/* <div style={{position: 'relative'}}> */}
                {/* use if you get a logo */}
                {/* <div class="content">
                    <div class="container">
                        <div class="row jumbotron bg-transparent">
                            <div class="col-md-6 text-center mt-5">
                                <img src={HomePageLogo} alt="inctfj-logo" height="100" width="200" />
                            </div>
                            <div class="col-md-6 text-center mt-5">
                                <h2>ARCHIVE</h2>
                                <h3>Learn | Hack | Repeat</h3>
                                <p class="lead">By team bi0s</p>

                                {buttons}

                            </div>
                        </div>
                    </div>
                </div> */}
                <div className="content">

                <div className="container">
                            <div className="jumbotron bg-transparent text-center mt-5">
                                <h2>Interal CTF</h2>
                                <h3>Learn | Hack | Repeat</h3>
                                <p className="lead">For team bi0s</p>

                                {buttons}

                        </div>
                    </div>
                </div>
            </React.Fragment >
        );
    }
}

export default HomePage;
