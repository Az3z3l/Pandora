import React, { Component } from "react";
import { api } from  "../../js/api.js"
import Button from 'react-bootstrap/Button'
import Table from 'react-bootstrap/Table'
import InputGroup from 'react-bootstrap/InputGroup'


class AdminEditchallenge extends Component {



    constructor(props) {
        super(props);

        const { challid } = this.props.match.params

        this.state = {
            ID: challid,
            name: '',
            description: '',
            category: '',
            tags: '',
            value: '',
            flags: '',
            visibility: null,
            isLoading: true,
            unavailable: false,
            delVal: null,
        }

    }      

    async componentDidMount() {


        const query = `
        query onechall($in:String!){
            onechallenge(id:$in){
              name
              description
              category
              tags
              value
              flags
              visibility
              file
            }
          }
        `
        let data = {
            in : this.state.ID
        }
        const newLocal = await api("onechall", data, query);
        var response = newLocal.onechallenge;
        if (response === null){
            this.setState({
                unavailable: true,
            })
        } else {
            this.setState({

                name: response.name,
                description: response.description,
                category: response.category.join(" "),
                tags: response.tags.join(" "),
                value: response.value,
                flags: response.flags,
                visibility: response.visibility,
                file: response.file

            });

            if (this.state.file == null){
                this.setState({
                    file: [],
                })
            }
        }
        this.setState({
			isLoading: false
		});
    }

    handleSubmit =async  (event) => {
        event.preventDefault();

        var data = {
            ID : this.state.ID,
            name : this.state.name,
            description : this.state.description,
            category : this.state.category.split(" "),
            tags : this.state.tags.split(" "),
            value : this.state.value,
            flags : this.state.flags,
        }
        let query = `mutation editChall($in: edit_challenge_data){
            edit_challenge(input:$in)
          }`

        let req = await api("editChall", { "in":data }, query);
        alert(req.edit_challenge)
 
        
    }
    onFileChange = event => { 
        this.setState({ selectedFile: event.target.files[0] }); 
    }; 

    handleInputChange = (event) => {
        this.setState({ 
            [event.target.name]: event.target.value 
        })
    }

    fileDelete = async (file) => {
        let datasend = {
            ID: this.state.ID,
            name: file
        }
        let operationName = "delfile";
		let variables = {
			"in": datasend,
		};
		let query = `
		mutation delfile($in: delfile){
            deletefile(input: $in)
          }
		`
		var a = await api(operationName, variables, query)
        var msg = a.deletefile;
        if (msg === "OK"){
            let x = this.state.file
            var index = x.indexOf(file);
            if (index !== -1) {
                x.splice(index, 1);
            }
            this.setState({
                file: x
            })
        }
    }

    handleDelete = async (event) => {
        let z = (window.confirm('You sure you wanna delete this challenge? The associated solves and challenges will not be deleted'))
        if (!z){
            return
        }
        let operationName = "delchall";
		let variables = {
			"in": this.state.ID,
		};
		let query = `
		mutation delchall($in: String!){
            deletechallenge(id: $in)
          }
        `
        var a = await api(operationName, variables, query)
        var msg = a.deletechallenge;
        if (msg === "OK"){
            document.location="../";
        }
        else{
            alert(msg);
            document.location="../";
        }
    }

    handleVisiblity = async (event) => {
        let x = !this.state.visibility
        let cond = (x)?'public':'private'
        let z = (window.confirm('Do you wish to make the challenge '+cond+'?'))
        if (!z){
            return;
        }
        let operationName = "visi";
        let datasend = {
            ID: this.state.ID,
            visibility: x
        }
		let variables = {
			"in": datasend,
		};
		let query = `
		mutation visi($in: public){
            challvisibility(input: $in)
          }
        `
        var a = await api(operationName, variables, query)
        var msg = a.challvisibility;
        alert(msg)
        this.setState({
            visibility: x,
        })
    }


    onFileUpload = async (event) => {
        if (this.state.selectedFile === null){
            return
        }
        let url = "/api/challenge/add"
        const data = new FormData();
        console.log(this.state.ID)
        data.append('id', this.state.ID);
        data.append('file',this.state.selectedFile );
        var responz = await fetch(url, {
            "credentials": "include",
            "method": "POST",
            "mode": "cors",
            headers: {
                'Accept': 'application/json',
            },
            body: data,
       })
       let c = await responz.json();
       if (c.Status==="OK"){
           let x = this.state.file;
           x.push(this.state.selectedFile.name)
           this.setState({
               file: x
           })
       } else {
           alert(c.Status)
       }
    }

    render() {
        let i = 1;


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
							<h2>Edit challenge</h2>
						</div>
						<br />
						<h3>Challenge not available</h3>
					</div>
				</React.Fragment>
			);
        }


        return (
            <div class="container h-100">
                <div class="row align-items-center h-90">
                    <div class="col-md-5 mx-auto mt-3">
                        <h2>Edit challenge</h2>
                        <hr />
                        <form name="register-data" >
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
                                <Button className="btn btn-block" variant="success" onClick={this.handleSubmit}>Save Edits</Button>
                            </div>
                            <div class="form-group">
                                <Button className="btn btn-block" variant="warning" onClick={this.handleVisiblity}>Toggle Visibility (currently {this.state.visibility?'public':'private'})</Button>
                            </div>
                            <div class="form-group">
                                <Button className="btn btn-block" variant="danger" onClick={this.handleDelete}>Delete Challenge</Button>
                            </div>
                        </form>
                    </div>

                    <div class="col-md-5 mx-auto mt-3">
                        <br />
                        <h4>File Management</h4>
                        <hr />
                        <h5>Available Files</h5>
                        <Table striped bordered hover variant="dark">
                            <thead>
                                <tr>
                                <th>#</th>
                                <th>Files</th>
                                <th>Delete</th>
                                </tr>
                            </thead>
                            <tbody>
                                {this.state.file.map((row) => {
                                    return (<>
                                        <tr>
                                        <td>{i++}</td>
                                        <td>{row}</td>
                                        <td><Button variant="outline-danger" onClick={() => this.fileDelete(row)}>&#128465;</Button></td>
                                        </tr>
                                        </>);
                                })}
                            </tbody>
                            </Table>
                            <h5>Choose file to upload</h5>
                            <Button as="input" type="file" variant="dark" onChange={this.onFileChange} />
                            <hr />
                            <Button onClick={this.onFileUpload}> Upload File </Button> 
                    </div>


                </div>
            </div>
        )
    }
}

export default AdminEditchallenge;
