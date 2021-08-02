import React, { Component } from "react";
import { api } from  "../../js/api.js"

class Adminaddchallenge extends Component {

    constructor(props) {
        super(props);

        this.state = {
            name: '',
            description: '',
            category: '',
            tags: '',
            value: '',
            flags: '',
        }

    }

    handleSubmit =async  (event) => {
        event.preventDefault();
        if(!Number.isInteger(parseInt(this.state.value))){
            alert("Enter Integer for Value")
            return
        }
        var data = {
            name : this.state.name,
            description : this.state.description,
            category : this.state.category.split(" "),
            tags : this.state.tags.split(" "),
            value : this.state.value,
            flags : this.state.flags,
        }
        let query = `mutation addChall($in: add_challenge_data){
            add_challenge(input:$in)
          }`

        let req = await api("addChall", { "in":data }, query);
        alert(req.add_challenge)
        if (req.add_challenge === "OK"){
            this.setState = ({
                name: '',
                description: '',
                category: '',
                tags: '',
                value: '',
                flags: '',
            })
        }
        
    }

    handleInputChange = (event) => {
        this.setState({ 
            [event.target.name]: event.target.value 
        })
    }

    render() {
        return (
            <div class="container h-100">
                <div class="row align-items-center h-90">
                    <div class="col-md-5 mx-auto mt-5">
                        <h2>Add challenge</h2>
                        <hr />
                        <form name="register-data" onSubmit={this.handleSubmit}>
                            <div class="form-group">
                                <input name="name" type="text" class="form-control" placeholder="Name *" value={this.state.name} onChange={this.handleInputChange} required />
                            </div>
                            <div class="form-group">
                                <textarea name="description" type="text" class="form-control" placeholder="Description *" value={this.state.description} onChange={this.handleInputChange} required />
                            </div>
                            <div class="form-group">
                                <input name="category" type="text" class="form-control" placeholder="Category *" value={this.state.category} onChange={this.handleInputChange} required />
                            </div>
                            <div class="form-group">
                                <input name="tags" type="text" class="form-control" placeholder="Tags *" value={this.state.tags} onChange={this.handleInputChange} required />
                            </div>
                            <div class="form-group">
                                <input name="value" type="text" class="form-control" placeholder="Value *" value={this.state.value} onChange={this.handleInputChange} required />
                            </div>
                            <div class="form-group">
                                <input name="flags" type="text" class="form-control" placeholder="Flag *" value={this.state.flags} onChange={this.handleInputChange} required />
                            </div>
                            
                            <div class="form-group">
                                <input type="submit" className="btn btn-success btn-block" value="Add challenge" required/>
                            </div>
                            
                        </form>
                    </div>
                </div>
            </div>
        )
    }
}

export default Adminaddchallenge;
