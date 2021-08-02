import React, { Component } from "react";

class Adminviewuser extends Component {
    render() {
        return (
            <React.Fragment>
                <div class="container">
                    <div class="row align-items-center h-100">
                        <div class="col-md-5 mx-auto">
                            <div class="bg-transparent mt-5">
                                <h2>View User</h2>
                                <hr />
                                <form name="signin">
                                    <div class="form-group">
                                        <input name="username" type="text" class="form-control" placeholder="Enter Username *"
                                            value="" required />
                                    </div>
                                    
                                    <div class="form-group">
                                        <input type="button" class="btn btn-success btn-block" value="Search" required
                                            onclick="login();" />
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

export default Adminviewuser;
