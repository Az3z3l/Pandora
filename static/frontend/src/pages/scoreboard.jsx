import React, { Component } from "react";
import { api } from  "../js/api.js"

class ScoreBoard extends Component {
	constructor(props) {
		super(props);
		this.state = {
			scoreboard: [],
			isLoading: true
		}

	}
	async componentDidMount() {
		// var query = `
		// query score{
		// 	scoreboard{
		// 		username
		// 		score
		// 	}
		// }`;

		// let a = await api("score", {}, query)
		
		var response = await fetch("/api/scoreboard", {
			headers: {
				"content-type": "application/json",
			},
			"method": "GET",
		});
		let x = ""
		if (response.ok){
			x = await response.json();
		} else {
			localStorage.removeItem('user');
			document.location='/login'
		}


		let res = x.scoreboard
		this.setState({
			scoreboard: res
		});
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
		let i = 1;

		return (
			<React.Fragment>
				<div className="container">
					<div className="col-md-12">
						<div className="row mt-5">
							<h2>Scoreboard</h2>
						</div>
						<div className="row mt-4">
						</div>
						<div className="row">
							<table className="table mt-4 Table">
								<thead>
									<tr>
										<th scope="col">#</th>
										<th scope="col">Username</th>
										<th scope="col">Score</th>
									</tr>
								</thead>
								<tbody>
									{this.state.scoreboard.map((row) => {
										return (
											<tr key={row.username}>
												<td>{i++}</td>
												<td>{row.username}</td>
												<td>{row.score}</td>
											</tr>
										);
									})}
								</tbody>
							</table>
						</div>
					</div>
				</div>
			</React.Fragment>
		)
	}
}

export default ScoreBoard;