import React, { Component } from "react";
import  {api} from "../../js/api.js"
import { Link } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.css";
import Table from 'react-bootstrap/Table'

class Challenges extends Component {

	constructor(props) {
		super(props);
		this.state = {
			challenge: [],
			isLoading: true,
			unavailable: false
		}
	}

	openModal = id => {
		this.setState({ openedModal: id });
	};
	closeModal = () => {
		this.setState({ openedModal: null });
	};

	async componentDidMount() {
		var query = `
		query challs{
			challenge {
			  ID
			  name
			  category
			  solves
			  }
		  }		  
		  `;

		// var datta = await api("challs", {},query)

		let a = await api("challs", {}, query)

		let datum = a.challenge;		

		if (datum == null){
			this.setState({
				unavailable: true
			})
		}

		this.setState({
			challenge: datum
		})

		this.setState({
			isLoading: false
		});

	}

	// openModal = () => this.setState({ isOpen: true });
	// closeModal = () => this.setState({ isOpen: false });


	render() {
		if (this.state.isLoading) {
			return (
				<React.Fragment>
					<div class="spinner"></div>
				</React.Fragment>
			);
		}
		let i = 1;

		if (this.state.unavailable){
			return (
				<React.Fragment>
					<div className="container">
						<div className="row mt-5">
							<h2>Challenges</h2>
						</div>
						<br />
						<h3>No Challenge available</h3>
					</div>
				</React.Fragment>
			);
		}

		return (
			<React.Fragment>
				<div className="container">
					<div className="row mt-5">
						<h2>Challenge Index</h2>
					</div>
					<br />
					<div className="row">
						<div className="col-md-8">
						
						<Table striped bordered hover responsive variant="dark">
							<thead>
								<tr>
								<th>#</th>
								<th>Name</th>
								<th>Category</th>
								<th>Solves</th>
								<th>Edit</th>
								</tr>
							</thead>
							<tbody>
							{this.state.challenge.map((row) => {
							return (<>
								<tr>
								<td>{i++}</td>
								<td>{row.name}</td>
								<td>{row.category.join(" ")}</td>
								<td>{row.solves}</td>
								<td><Link to={`edit/${row.ID}`}>edit</Link></td>
								</tr>
								</>);
						})}

							</tbody>
						</Table>
						</div>
					</div>
				</div>
			</React.Fragment>
		);
	}
}

export default Challenges;
