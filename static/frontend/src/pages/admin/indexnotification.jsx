import React, { Component } from "react";
import  {api} from "../../js/api.js"
import { Link } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.css";
import Table from 'react-bootstrap/Table'

class Adminindexnotification extends Component {

	constructor(props) {
		super(props);
		this.state = {
			notifications: [],
			isLoading: true,
			unavailable: false
		}
	}

	async componentDidMount() {
		var query = `
		query notifications{
            notify{
                ID
                name
                description	
                timestamp
            }
        }		  
		  `;

		let a = await api("notifications", {}, query)

		let datum = a.notify;		

		if (datum == null){
			this.setState({
				unavailable: true
			})
		}

		this.setState({
			notifications: datum
		})

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

		if (this.state.unavailable){
			return (
				<React.Fragment>
					<div className="container">
						<div className="row mt-5">
							<h2>Notifications</h2>
						</div>
						<br />
						<h3>No Notifications available</h3>
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
								<th>Time of Release</th>
								<th>Edit</th>
								</tr>
							</thead>
							<tbody>
							{this.state.notifications.map((row) => {
							return (<>
								<tr>
								<td>{i++}</td>
								<td>{row.name}</td>
								<td>{row.timestamp}</td>
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

export default Adminindexnotification;
