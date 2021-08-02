import React, { Component } from "react";
import { api } from  "../js/api.js"
import * as moment from 'moment'
import Moment from "react-moment";
import { Redirect } from 'react-router-dom'

class Profile extends Component {
	constructor(props) {
		super(props);
		this.state = {
			username: '',
			solves: [],
			score: 0,
			auth: true,
			isLoading: true
		}
	}

	async componentDidMount() {
		var query = `
		query datau{
			userid {
			  username
			  score 
			  solved {
				Timestamp
				challenge {
				  name
				  value
				  category
				}
				
			  }
			}
		  }
        `;
		let response = await api("datau", {}, query)
		
		if (response.userid.solved!=null){
			this.setState({
				username: [response.userid.username],
				solves: [response.userid.solved][0],
				score: [response.userid.score]
			})
		} else {
			this.setState({
				username: [response.userid.username],
				solves: [],
				score: [response.userid.score]
			})
		}

		this.setState({
			isLoading: false
		});
	}

	render() {
		if (this.state.isLoading) {
			return (
				<React.Fragment>
					<div class="spinner"></div>
				</React.Fragment>
			);
		}
		let i = 0;
		let rows;
		if (this.state.solves != null) {
			rows = this.state.solves.map((row) => {
				return (
					<tr key={i++}>
						<td>{i}</td>
						<td>{row.challenge.name}</td>
						<td>{row.challenge.category.join(',')}</td>
						<td>{row.challenge.value}</td>
						<td><Moment format="MMM	DD YYYY h:m:s">{row.Timestamp}</Moment></td>
					</tr>
				)
			})
		}
		else {
			rows =
				(<tr>
					<td>No solves</td>
				</tr>)
		}

		return (
			<React.Fragment>
				<main>
					<div className="container h-100">
						<div className="row align-items-center h-90">
							<div className="col-sm-12">
								<div className="jumbotron bg-transparent">
									<h2>{this.state.username} - {this.state.score} pts</h2>
									<hr />
									<h3>Solves</h3>
									<div className="row">
										<table className="table mt-4 Table">
											<thead>
												<tr>
													<th scope="col">#</th>
													<th scope="col">Challenge Name</th>
													<th scope="col">Category</th>
													<th scope="col">Points</th>
													<th scope="col">Solved on</th>
												</tr>
											</thead>
											<tbody>
												{
													rows
												}

											</tbody>
										</table>
									</div>
								</div>
							</div>
						</div>
					</div>
				</main>
			</React.Fragment >
		);
	}
}

export default Profile;
