package bolt

/*

	private double EN1993_1_8_TABLE_3_3_e1_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*1.2;
	}

	private double EN1993_1_8_TABLE_3_3_e2_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*1.2;
	}

	private double EN1993_1_8_TABLE_3_3_e3_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*1.5;
	}

	private double EN1993_1_8_TABLE_3_3_e4_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*1.5;
	}

	private double EN1993_1_8_TABLE_3_3_p1_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*2.2;
	}

	private double EN1993_1_8_TABLE_3_3_p2_min(int DiameterBolt)
	{
	    return BOLT_Do(DiameterBolt)*2.4;
	}

	String Output()
	{
		String out = new String();
		out = "";
	    out += "Bolt: " + PrintfDia(DiameterBolt) + "\n";
	    out += "Class of bolt: " + PrintfBS(BS) + "\n";
	    out += "Diameter of hole: " + DiameterHole + " mm\n";
	    out += "Partial safety factor for joint:\n";
	    out += String.format("gammaM2 = %.2f\n",gamma_M2);
	    out += "\n";
	    out += "The gross cross-section area of bolt:\n";
	    out += String.format("A  = %.1f sq.mm\n",A );
	    out += "The tensile stress area of the bolt:\n";
	    out += String.format("As = %.1f sq.mm\n",As);
	    out += "\n";
	    out += "Nominal value of the yield strength(table 3.1 EN1993-1-8):\n";
	    out += String.format("Fyb  = %.1f MPA\n",EN1993_1_8_TABLE_3_1_Fyb(BS) );
	    out += "Nominal value of the ultimate tensile strength(table 3.1 EN1993-1-8):\n";
	    out += String.format("Fub  = %.1f MPA\n",EN1993_1_8_TABLE_3_1_Fub(BS) );
	    out += "\n";
	    out += "Shear resistance per shear plane(table 3.4 EN1993-1-8):\n";
	    out += String.format("Fv,Rd = %.1fkN\n",F_v_Rd*1e-3);
	    out += "Tension resistance(table 3.4 EN1993-1-8):\n";
	    out += String.format("Ft,Rd = %.1fkN\n",F_t_Rd*1e-3);
	    out += "\n";
	    out += "Minimal spacing, end and edge distances(table 3.3 EN1993-1-8):\n";
	    out += String.format("End  distance e1  = %.1f mm\n",EN1993_1_8_TABLE_3_3_e1_min(DiameterBolt) );
	    out += String.format("Edge distance e2  = %.1f mm\n",EN1993_1_8_TABLE_3_3_e2_min(DiameterBolt) );
	    out += String.format("Distance e3 in slotted holes = %.1f mm\n",EN1993_1_8_TABLE_3_3_e3_min(DiameterBolt) );
	    out += String.format("Distance e4 in slotted holes = %.1f mm\n",EN1993_1_8_TABLE_3_3_e4_min(DiameterBolt) );
	    out += String.format("Spacing p1 = %.1f mm\n",EN1993_1_8_TABLE_3_3_p1_min(DiameterBolt) );
	    out += String.format("Spacing p2 = %.1f mm\n",EN1993_1_8_TABLE_3_3_p2_min(DiameterBolt) );

	    // *
	    if(B_p_Rd < 1e20)
	    	out += "B_p_Rd = %.1f kN\n",B_p_Rd*1e-3);* //
	    return out;
	};
};
*/
