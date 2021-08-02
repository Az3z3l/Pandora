import React, { Component } from "react";
import { api } from  "../../js/api.js"
import Button from 'react-bootstrap/Button'
import Form from 'react-bootstrap/Form'
import Radio from '@material-ui/core/Radio';
import RadioGroup from '@material-ui/core/RadioGroup';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import FormControl from '@material-ui/core/FormControl';
import FormLabel from '@material-ui/core/FormLabel';



class Adminctfmanagement extends Component {

    constructor(props) {
        super(props);


        this.state = {
            status: '',
            details: '',
            unavailable: false,
            isLoading: true,
        }

        // const [value, setValue] = React.useState('female');

    }
    
    async componentDidMount() {


        const query = `
        query settings {
            frontendmanagement{
                status
                details
                scoreboardStats
          }
        }
        `
     
        const newLocal = await api("settings", {}, query);
        var response = newLocal.frontendmanagement;
        this.setState({
            status: response.status.toString(),
            details: response.details,
            scoreboardStats: response.scoreboardStats.toString(),
        });

        
        this.setState({
			isLoading: false
		});
    }

    handleInputChange = (event) => {
        this.setState({ 
            [event.target.name]: event.target.value 
        })
    }

    // handleChange = (event) => {
    //     setValue(event.target.value);
    //   };
  

    changeStat = async (event) =>{
        event.preventDefault();
        var data = {
            status: this.state.status,
            details : this.state.details,
            scoreboardStats: this.state.scoreboardStats
        }
        let query = `mutation setting($in :Manager){
            adminmanagement(input: $in)
          }`

        let req = await api("setting", { "in": data }, query);
        alert(req.adminmanagement)
    }  

    render() {
        if (this.state.isLoading) {
			return (
				<React.Fragment>
					<div class="spinner"></div>
				</React.Fragment>
			);
        }
     
        else{
            return (
                <div class="container h-100">
                    <div class="row align-items-center h-90">
                        <div class="col-md-5 mx-auto mt-5">
                            <h2>Challenge and Scoreboard settings</h2>
                            <hr />

                            <form name="register-data">
                                <FormControl>
                                <label><h4>Set CTF Status</h4></label>
                                    <RadioGroup aria-label="status" name="status" value={this.state.status} onChange={this.handleInputChange}>
                                        <FormControlLabel value="0" control={<Radio />} label="Start/Resume CTF" />
                                        <FormControlLabel value="1" control={<Radio />} label="Stop CTF - hide all challenges, no flag check " />
                                        <FormControlLabel value="2" control={<Radio />} label="Stop CTF - show challenges, flag check possible" />
                                        <FormControlLabel value="3" control={<Radio />} label="Stop CTF - show challenges, no flag check" />
                                    </RadioGroup>
                                </FormControl>
                                <div class="form-group">
                                    <label><h4>Why?</h4></label>
                                    <textarea name="details" type="text" class="form-control" placeholder="Why *" value={this.state.details} onChange={this.handleInputChange} required />
                                </div>
                                <FormControl>
                                <label><h4>Scoreboard visiblity</h4></label>
                                    <RadioGroup aria-label="status" name="scoreboardStats" value={this.state.scoreboardStats} onChange={this.handleInputChange}>
                                        <FormControlLabel value="0" control={<Radio />} label="Only LoggedIn Users" />
                                        <FormControlLabel value="1" control={<Radio />} label="Open to all" />
                                    </RadioGroup>
                                </FormControl>
                                <div class="form-group">
                                    <input type="button" onClick={this.changeStat} class="btn btn-success btn-block" value="Change Status" required/>
                                </div>
                                
                            </form>
                        </div>
                    </div>
                </div>
            )
        }
    }
}

export default Adminctfmanagement;
