import React, { Component } from "react";
import { Nav, NavItem, NavLink } from 'reactstrap';

class NotFound extends Component {
    render() {
        return (
            <React.Fragment>
            <div className="container">
					<div className="col-md-12">
						<div className="col mt-5">
							<h2>Page not found</h2>
                            <h4>r/100</h4>
						</div>
						<div className="row mt-4">
						</div>
						<div className="col-md-12">
                        
                            <Nav vertical className="text-center">
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=sLisEEwYZvw">I'll Be There For You</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=IYq1bsyvT5Q">The Last Play</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=8ALsR-oWKAk">Goodbye, Michael</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=kNw8V_Fkw28">I'm not crying, you are. </NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=3pCGU3bIPEc">Ramin Djawadi's Magic</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=MASPbl210_c">* laughs *</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=eTa1jHk1Lxc">The Unsung Heroes </NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=sv7d3TV7btA">The Ackerman Rage</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=G-bIgXS9tU8">I Talked To You For The First Time</NavLink>
                                </NavItem>
                                <NavItem>
                                    <NavLink href="https://www.youtube.com/watch?v=spCdFMnQ1Fk">Goodbye Beautiful</NavLink>
                                </NavItem>
                            </Nav>       

						</div>
					</div>
				</div>
            </React.Fragment>
        )
    }
}


export default NotFound;
