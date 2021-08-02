import React, { Component } from "react";
import { Link } from "react-router-dom";
import HomePageLogo from '../../images/d2.png';

class Admindashboard extends Component {
    render() {
        return (
            <React.Fragment>
                <div class="content">
                    <div class="container">
                        <div class="row jumbotron bg-transparent">
                            
                            <div class="col-md-25 text-center mt-5">
                                <h3>Admin Dashboard</h3>
                                <center>
                                <p class="lead">
                                    <Link className="btn btn-primary btn-lg" to={`/admin/challenges`}>
                                        View Challenge
                                    </Link>
                                    <Link className="btn btn-primary btn-lg" to={`/admin/challenges/add`}>
                                        Add Challenge
                                    </Link>
                                    <Link className="btn btn-primary btn-lg" to={`/admin/challenges/edit`}>
                                        Edit Challenge
                                    </Link>
                                    <Link className="btn btn-primary btn-lg" to={`/admin/users/view`}>
                                        User Details
                                    </Link>
                                    <Link className="btn btn-primary btn-lg" to={`/admin/notifications`}>
                                        Notifications
                                    </Link>

                                </p></center>
                            </div>
                        </div>
                    </div>
                </div>
            </React.Fragment >
        );
    }
}

export default Admindashboard;
