import React, { Component } from "react";
import { api } from  "../js/api.js"
import Badge from 'react-bootstrap/Badge'
import "bootstrap/dist/css/bootstrap.css";
import { Button, Modal, ModalHeader, ModalBody, ModalFooter } from "reactstrap";
import ReactMarkdown from "react-markdown";
import Tab from 'react-bootstrap/Tab'
import Nav from 'react-bootstrap/Nav'
import TabContainer from 'react-bootstrap/TabContainer'
import TabContent from 'react-bootstrap/TabContent'
import TabPane from 'react-bootstrap/TabPane'

class Challenges extends Component {

	constructor(props) {
		super(props);
		this.state = {
			challenge: [],
			flag: '',
			msg: '',
			msg_type: '',
			openedModal: null,
			solved: [],
			auth: true,
			isLoading: true,
			status: '',
			details: '',
			getChall: true,
			print: false,
		}
	}

	openModal = id => {
		this.setState({ openedModal: id });
	};
	closeModal = () => {
		this.setState({ openedModal: null });
	};

	  

	async componentDidMount() {
		let query = `
        query settings {
            frontendmanagement{
                status
                details
          }
        }
        `
     
        let newLocal = await api("settings", {}, query);
        var response = newLocal.frontendmanagement;
        this.setState({
            status: response.status.toString(),
            details: response.details,
        });

		let stat= this.state.status
		if (stat !== null){
			if (stat === '1'){
				this.setState({
					getChall: false,
				})
			}
			if (stat !== '0'){
				this.setState({
					print: true,
				})
			}
		}

		if (!this.state.getChall){
			this.setState({
				isLoading: false
			});
			return
		}


		var query1 = `
		query mydata{
			userid {
			  solved {
				challenge {
				  ID
				}
			  }
			}
		  }`;

		let b =await api("mydata", {}, query1)
		let sol = []
		if (b.userid.solved!==null){
			for (var i = 0; i < b.userid.solved.length; i++) {
				sol.push(b.userid.solved[i].challenge.ID)
			}
		}
		this.setState({
			solved: sol
		})

		query = `
		query challs{
			challenge {
			  ID
			  name
			  value
			  description
			  category
			  tags
			  file
			  solves			  
			  }
		  }		  
		  `;

		// var datta = await api("challs", {},query)

		let a = await api("challs", {}, query)

		let datum = a.challenge;
		if (datum === null){
			this.setState({
				print: true,
			})
			if (this.setState.details === null){
				this.setState({
					details: "CTF has not started or no Challenges are available",
				})
			}
		}
		let group=[]
		if (datum !== null || datum !== undefined){
				group = datum.reduce((r, a) => {
				r[a.category[0]] = [...r[a.category[0]] || [], a];
				return r;
			}, {});
		}
		this.setState({
			challenge: group
		})

		this.setState({
			isLoading: false
		});

	}

	handleSubmit = async (event) => {
		event.preventDefault();

		let datasend = this.state.flag
		let operationName = "flaggy";
		let variables = {
			"in": datasend,
		};
		let query = `
		mutation flaggy($in: String!){
			flag_submit(input:$in)
		}
		`
		var a = await api(operationName, variables, query)
		var msg = a.flag_submit;

		this.setState({
			msg: msg
		})
		if (this.state.msg === "Check your flag again") {
			var type = 'alert alert-danger'
			this.setState({
				msg_type: type
			})
		}
		else if (msg === "Congratulations on the solve") {
			var type = 'alert alert-success'
			this.setState({
				msg_type: type
			})
		}
		else if (msg === "Already solved this Challenge") {
			var type = 'alert alert-info'
			this.setState({
				msg_type: type
			})
		}
		else {
			var type = 'alert alert-info'
			this.setState({
				msg_type: type
			})
		}
		setTimeout(function () {
			this.setState({
				msg: '',
				msg_type: ''
			});
		}.bind(this), 20000);

	};

	onChange = (e) => {
		this.setState({
			flag: e.target.value
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

		if (!this.state.getChall){
			return (
			<div className="container">
					<div className="row mt-5">
						<h2>Challenges</h2>
					</div>
					<br />
					<h3>{this.state.details}</h3>
			</div>
			);
		}

		return (
			<React.Fragment>
				<div className="container">
					<div className="row mt-5">
						<h2>Challenges</h2>
					</div>
					<br />
					{this.state.print === true ? this.state.details: '' }
					<div className="row">
						<div className="col-md-8">
						<Tab.Container id="challs" defaultActiveKey="sanity">

							<Nav variant="tabs" className="flex-row">
								{Object.entries(this.state.challenge).map(([key, value]) => {
									return (<>
										<Nav.Item>
											<Nav.Link eventKey={key}>{key}</Nav.Link>
										</Nav.Item>
								</>)})}
							</Nav>
							<hr />

							<Tab.Content>
								{Object.entries(this.state.challenge).map(([key, value]) => {
									return (<>
									<Tab.Pane eventKey={key}>
										{/* <div key={key} className="row mt-5">
											<h3>{key}</h3>
										</div> */}
										<div className="row">
											{value.map((row) => {
												return (<>
													<div key={row.ID} className="col-sm-4 py-2">
														<div onClick={() => this.openModal(row.ID)} className={this.state.solved.indexOf(row.ID) !== -1 ? 'card h-100 w-200 bg-success' : 'card h-100 w-200 bg-dark'}>
															<div className="card-body text-center">
																<h6 className="card-title h-50 text-center card">{row.name}</h6>
																<h5 className="card-text h-50 text-center">{row.value}</h5>
																<input type="hidden" name="key" value={row.ID} />
																<div className="challengefont">
																	<Modal
																		isOpen={this.state.openedModal === row.ID}
																		toggle={this.closeModal}
																		className="challengefont"
																	>
																		<ModalHeader>{row.name}</ModalHeader>
																		<ModalBody>
																			<h5><ReactMarkdown source={row.description} /></h5>
																			<br />
																		solves : {row.solves}
																			<br />
																		categories : <h5>
																				{row.category.map((cats) => {
																					return (<>
																						{' '}<Badge pill variant="danger">
																							{cats}
																						</Badge>
																					</>);
																				}
																				)}</h5>
																			<br />
																			{(row.file !== null || row.file !== undefined) &&        
																			<h5>
																				{/* File: <a href={`/files/${this.state.challid}/${this.state.challenge.file}`} download>{this.state.challenge.file}</a> */}
																				File: 
																				{row.file.map((cats) => {
																					return (<>
																						{' '}
																						<a href={`/files/${row.ID}/${cats}`} download>{cats}</a>
																					</>);
																				}
																				)}
																			</h5>  
																		}
																		tags : <h5>
																				{row.tags.map((taggy) => {
																					return (<>
																						{' '}<Badge variant="light">
																							{taggy}
																							</Badge>
																					</>);
																				}
																				)} </h5>

																		</ModalBody>
																		{/* <ModalFooter>
																			<Button color="primary" onClick={this.closeModal}>
																				Close
																		</Button>
																		</ModalFooter> */}
																	</Modal>
																</div>
															</div>
														</div>
													</div>
												</>);
											})}
										</div>
									</Tab.Pane>
									</>);
								})}
							</Tab.Content>
						</Tab.Container>
						</div>
						<div className="col-md-4" id="btext1">
							<div className="row mt-5">
								<div className="col-md-12">
									<div className="card mt-4 bg-dark">
										<form name="flag-submit" onSubmit={this.handleSubmit}>
											<div className="card-body">
												<h6 className="card-title">Flag submission</h6>
												<div className="form-group">
													<input type="text" className="form-control" placeholder="Enter flag *" required onChange={this.onChange} />
												</div>
												<input type="submit" className="btn btn-success btn-block" value="Submit Flag" />
												<br />
												<div className={this.state.msg_type} onChange={this.timer}>
													<strong>{this.state.msg}</strong>
												</div>
											</div>
										</form>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>
			</React.Fragment>
		);
	}
}

export default Challenges;
