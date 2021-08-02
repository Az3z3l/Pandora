import React, { Component } from "react";
import { validateContact, validateEmail } from '../js/helper.js';
import { Link } from "react-router-dom";


class Register extends Component {

    componentDidMount() {
        document.location="https://junior.inctf.in/register"
    }

     render() {
        return (
            <React.Fragment>
    {/* 
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
    */}
             </React.Fragment>
        )
    }
}

export default Register;