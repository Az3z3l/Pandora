import React, { Component } from "react";
import { Link } from "react-router-dom";
import logo from '../images/logo.png';
import { Button } from "react-bootstrap"
import { Navbar, Nav, NavDropdown } from 'react-bootstrap'
import '../css/style.css'

export const Header = class Header extends Component {

    logout = () => {
        localStorage.clear();
        document.cookie = "auth= ; expires = Thu, 01 Jan 1970 00:00:00 GMT"
        document.location='/'
    }
    

    render() {
        const admin = localStorage.getItem('Jedi');
        const user = localStorage.getItem('user')
      
        if (admin){
            return (
                <React.Fragment>
                    <div id="particle-js"></div>
                    <header className="header">
                        <div className="container">
                            <nav className="navbar navbar-expand-lg">
                                <a className="navbar-brand" href="#">
                                    <img src={logo} className="main-logo" />
                                </a>
                                <button className="navbar-toggler" type="button" data-toggle="collapse"
                                    data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false"
                                    aria-label="Toggle navigation">
                                    <span className="navbar-toggler-icon"></span>
                                </button>
    
                                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                                    <ul className="navbar-nav mr-auto">
                                        <li className="nav-item active">
                                            <Link className="nav-link" to={`/`}>Home</Link>
                                        </li>
                                        <li className="nav-item">
                                            <Link className="nav-link" to={`/scoreboard`}>Scoreboard</Link>
                                        </li>
                                        <li>
                                        <NavDropdown title="Notifications" id="collasible-nav-dropdown">
                                            <NavDropdown.Item variant="dark"><Link to={`/notifications`}>View Existing</Link></NavDropdown.Item>
                                            <NavDropdown.Item variant="dark"><Link to={`/admin/notifications/add`}>Add new notification</Link></NavDropdown.Item>
                                            <NavDropdown.Item variant="dark"><Link to={`/admin/notifications/edit`}>Edit notification</Link></NavDropdown.Item>
                                        </NavDropdown>
                                        </li>
                                        <li>
                                        <NavDropdown title="Challenge" id="collasible-nav-dropdown">
                                            <NavDropdown.Item variant="dark"><Link to={`/challenges`}>User View</Link></NavDropdown.Item>
                                            <NavDropdown.Item variant="dark"><Link to={`/admin/challenges`}>All Challenges</Link></NavDropdown.Item>
                                            <NavDropdown.Item variant="dark"><Link to={`/admin/challenges/add`}>Add Challenge</Link></NavDropdown.Item>
                                            <NavDropdown.Item variant="dark"><Link to={`/admin/challenges/edit`}>Edit Challenge</Link></NavDropdown.Item>
                                        </NavDropdown>
                                        </li>
                                        <li>
                                        <NavDropdown title="Management" id="collasible-nav-dropdown">
                                            <NavDropdown.Item variant="dark"><Link to={`/admin/ctf/manage`}>User View</Link></NavDropdown.Item>
                                        </NavDropdown>
                                        </li>
                                    </ul>
                                    <ul className="navbar-nav p-2 login-btn">
                                        <Button variant="secondary"  onClick={this.logout}   >Logout</Button>
                                    </ul>
                                </div>
                            </nav>
                        </div>
                    </header>
                    <div className="spectrum"></div>
                </React.Fragment >
            );}
        else if (user){
        return (
            <React.Fragment>
                <div id="particle-js"></div>
                <header className="header">
                    <div className="container">
                        <nav className="navbar navbar-expand-lg">
                            <a className="navbar-brand" href="#">
                                <img src={logo} className="main-logo" />
                            </a>
                            <button className="navbar-toggler" type="button" data-toggle="collapse"
                                data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false"
                                aria-label="Toggle navigation">
                                <span className="navbar-toggler-icon"></span>
                            </button>

                            <div className="collapse navbar-collapse" id="navbarSupportedContent">
                                <ul className="navbar-nav mr-auto">
                                    <li className="nav-item active">
                                        <Link className="nav-link" to={`/`}>Home</Link>
                                    </li>
                                    <li className="nav-item">
                                        <a className="nav-link" href="http://wiki.bi0s.in/">Wiki</a>
                                    </li>

                                    <li className="nav-item">
                                        <Link className="nav-link" to={`/scoreboard`}>Scoreboard</Link>
                                    </li>
                                    <li className="nav-item">
                                        <Link className="nav-link" to={`/challenges`}>Challenges</Link>
                                    </li>
                                    <li className="nav-item">
                                            <Link className="nav-link" to={`/notifications`}>Notifications</Link>
                                    </li>
                                    <li className="nav-item">
                                            <Link className="nav-link" to={`/profile`}>Profile</Link>
                                    </li>
                                    <li className="nav-item">
                                            <Link className="nav-link" to={`/settings`}>Settings</Link>
                                    </li>
                                </ul>
                                <ul className="navbar-nav p-2 login-btn">
                                        <Button variant="secondary"  onClick={this.logout}   >Logout</Button>
                                </ul>
                            </div>
                        </nav>
                    </div>
                </header>
                <div className="spectrum"></div>
            </React.Fragment >
        );}
        else{
            return (
                <React.Fragment>
                    <div id="particle-js"></div>
                    <header className="header">
                        <div className="container">
                            <nav className="navbar navbar-expand-lg">
                                <a className="navbar-brand" href="#">
                                    <img src={logo} className="main-logo" />
                                </a>
                                <button className="navbar-toggler" type="button" data-toggle="collapse"
                                    data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false"
                                    aria-label="Toggle navigation">
                                    <span className="navbar-toggler-icon"></span>
                                </button>
    
                                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                                    <ul className="navbar-nav mr-auto">
                                        <li className="nav-item active">
                                            <Link className="nav-link" to={`/`}>Home</Link>
                                        </li>
                                        <li className="nav-item">
                                            <a className="nav-link" href={`/scoreboard`}>Scoreboard</a>
                                        </li>
                                    </ul>
                                    <ul className="navbar-nav p-2 login-btn">
                                        <Link className="btn btn-primary btn-lg" to={`/login`}>
                                            Login
                                        </Link>
                                       
                                        <Link className="btn btn-primary btn-lg" to={`/register`}>
                                            Register
                                        </Link>
                                    </ul>
                                </div>
                            </nav>
                        </div>
                    </header>
                    <div className="spectrum"></div>
                </React.Fragment >
            );
        }
    }
}


export const Footer = class Footer extends Component {
    render() {
        return (
            <React.Fragment>
                <footer>
                    <div className="container-fluid mt-5 text-center">
                        <p>Powered by <a href="http://www.bi0s.in">team bi0s</a></p>
                    </div>
                </footer>
            </React.Fragment>
        );
    }
}
